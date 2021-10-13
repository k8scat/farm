package exchange

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/pkg/errors"

	"github.com/reactivex/rxgo/v2"
)

// Subscriber 同步后的数据进行整理后拆成单独的event，并推送到订阅者
// 订阅者应该可能至少只有一个实例来接收farm产生的事件
type Subscriber interface {
	// 是否开启状态
	IsEnable() bool
	// Label 订阅者的标签
	Label() string
	// Actions 订阅者的事件类型
	Actions() []Action
	// Handle 处理event
	Handle(event *Event) error
	// LastOffset 订阅者的最后一个offset
	// 下面的2个方法可以通过组合 offset.go 中的结构体实现
	LastOffset() uint64
	// SetOffset 设置订阅者的最后一个offset
	SetOffset(uint64) error
}

type OrderlyMQ interface {
	// Run 运行mq
	Run() error
	// Register 注册订阅者
	Register(...Subscriber)
	// Push 将事件推给MQ
	Push(*Event) error
	// Pipe 从MQ中获取事件
	Pipe() rxgo.Observable
}

// Exchange 事件管理、推送中心
// 订阅者管理中心
// 将消息推送mq（mq将根据订阅者的情况，选择性的推送到订阅者）
type Exchange struct {
	ctx context.Context
	// 有序的队列
	mq OrderlyMQ
}

func New(ctx context.Context, mq OrderlyMQ) *Exchange {
	e := &Exchange{}
	e.ctx = ctx
	e.mq = mq
	return e
}

func (e *Exchange) Start() {
	go e.mq.Run()
	go e.run()
}

func (e *Exchange) AddSubscriber(sub Subscriber) {
	e.mq.Register(sub)
}

func (e *Exchange) Push(event *Event) error {
	// push event
	return e.mq.Push(event)
}

func (e *Exchange) run() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Exchange run: %+v - %s\n", e, debug.Stack())
		}
	}()

	for {
		select {
		case <-e.ctx.Done():
			log.Printf("Exchange exit.\n")
			return
		case eventV := <-e.mq.Pipe().Observe():
			event := eventV.V.(*PipeEvent)
			err := event.Wait(func(e *Event) error {
				err := event.affectedSubscriber.Handle(e)
				return errors.WithStack(err)
			})
			if err != nil {
				log.Printf("Exchange awarding event was err: %+v, this event '%d' had retry '%d' times.\n",
					err, event.event.Offset, event.retry)
			}
		}
	}
}
