package exchange

import (
	"log"
	"runtime/debug"
	"sync"
	"time"

	"github.com/reactivex/rxgo/v2"
)

// MQDB is a message queue is based on database
// It's needs 'farm_event' and 'farm_subscriber' tables

const (
	RETRY_COUNT    = 3
	RETRY_INTERVAL = 3 * time.Second
)

// MQDB 基于数据库的MQ
// 如果对于持久性要求不高的话，可以考虑使用 Redis Stream 来实现 MQ
//
type MQDB struct {
	pipe chan *PipeEvent

	// 当需要从数据库读取event时，该channel会被触发
	act chan rxgo.Item

	subWG       sync.Mutex
	subscribers map[string]Subscriber
}

func NewMQDB() *MQDB {
	mq := new(MQDB)
	mq.subscribers = make(map[string]Subscriber)
	mq.pipe = make(chan *PipeEvent, 1)
	mq.act = make(chan rxgo.Item, 128)
	return mq
}

func (d *MQDB) Run() error {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("MQDB: %s\n", debug.Stack())
		}
	}()
	// TODO 从数据库中读取 Subscribers 的offset并初始化
	// TODO 触发从数据库读取Event的工作
	//	 TODO 根据 subscribers 从数据库中以流的方式读取Event，并推送rxgo.Pipe，超时未处理或重试n次后未被处理的event将被放弃
	select {
	case <-rxgo.FromChannel(d.act).Last().Observe():

	}

	return nil
}

func (d *MQDB) Push(e *Event) error {
	// TODO 推送到数据库
	//
	// 触发从数据库读取Event的工作
	d.act <- rxgo.Item{}
	return nil
}

func (d *MQDB) Subscribers(subs ...Subscriber) {
	d.subWG.Lock()
	defer d.subWG.Unlock()

	for _, sub := range subs {
		d.subscribers[sub.Label()] = sub
	}
}

func (d *MQDB) Pipe() *PipeEvent {
	return <-d.pipe
}

func (d *MQDB) ReadStreamOfEvent() {
	for _, sub := range d.subscribers {
		sub.LastOffset()
		// TODO
	}
}
