package exchange

import (
	"context"
	"encoding/json"
	"log"
	"runtime/debug"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/molizz/farm/model/db"
	eventModel "github.com/molizz/farm/model/event"
	"github.com/pkg/errors"
	"github.com/reactivex/rxgo/v2"
)

// MQDB 基于数据库的MQ
// 如果对于持久性要求不高的话，可以考虑使用 Redis Stream 来实现 MQ
//
type MQDB struct {
	ctx context.Context

	cfg *MQConfig

	pipeSource chan rxgo.Item
	pipe       rxgo.Observable

	// 当需要从数据库读取event时，该channel会被触发
	act chan rxgo.Item

	subWG             sync.Mutex
	subscribers       map[string]Subscriber
	subscriberProcess *Process // Subscriber的进度管理
}

func NewMQDB(ctx context.Context, cfg *MQConfig) *MQDB {
	mq := new(MQDB)
	mq.ctx = ctx
	mq.cfg = cfg
	mq.pipeSource = make(chan rxgo.Item, 128)
	mq.pipe = rxgo.FromChannel(mq.pipeSource)
	mq.act = make(chan rxgo.Item, 128)
	mq.subscribers = make(map[string]Subscriber)
	mq.subscriberProcess = NewProcess(1)
	return mq
}

func (d *MQDB) Run() error {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("MQDB: %s\n", debug.Stack())
		}
	}()

	rg := rxgo.FromChannel(d.act)

	for {
		select {
		case <-d.ctx.Done():
			return nil
		case eventV := <-rg.Observe():
			event := eventV.V.(*Event)
			d.readEvents(event)
		}
	}
}

func (d *MQDB) Push(e *Event) error {
	// 推送到数据库
	dbErr := db.Transact(func(tx sqlx.Ext) error {
		err := eventModel.New(tx).Create(e.ToModel())
		return errors.WithStack(err)
	})

	// 触发从数据库读取Event
	if dbErr == nil {
		d.act <- rxgo.Item{V: e}
	}
	return dbErr
}

func (d *MQDB) Register(subs ...Subscriber) {
	d.subWG.Lock()
	defer d.subWG.Unlock()

	for _, sub := range subs {
		d.subscribers[sub.Label()] = sub
	}
}

func (d *MQDB) Pipe() rxgo.Observable {
	return d.pipe
}

func (d *MQDB) readEvents(event *Event) {
	for _, sub := range d.subscribers {
		if !sub.IsEnable() {
			log.Printf("Subscriber '%s' is disable, skip.\n", sub.Label())
			continue
		}

		actionExist := false
		for _, ac := range sub.Actions() {
			if ac == event.Action {
				actionExist = true
			}
		}
		if !actionExist {
			log.Printf("Subscriber '%s' actions '%s', Does not match for event actions for '%s', skip.\n",
				sub.Label(), sub.Actions(), event.Action)
			continue
		}

		go d.doReadEvents(sub, event)
	}
}

func (d *MQDB) doReadEvents(sub Subscriber, event *Event) {
	// 同一时刻，同一个Subscriber，只允许一个实例运行
	defer d.subscriberProcess.Start(sub.Label()).Wait()

	namespace := event.Context.Namespace

	err := eventModel.New(db.GetDB()).ListByNamespaceOnStream(namespace, sub.LastOffset(),
		func(dbEvent *eventModel.Event) error {
			newEvent := new(Event)
			err := json.Unmarshal([]byte(dbEvent.Payload), newEvent)
			if err != nil {
				return errors.WithStack(err)
			}

			d.pipeSource <- rxgo.Item{
				V: &PipeEvent{
					event:              newEvent,
					affectedSubscriber: sub,
					observable:         d.pipe,
				},
			}
			return nil
		})
	if err != nil {
		log.Printf("Reading events was err: %+v\n", err)
	}
}
