package exchange

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/molizz/farm/thirdparty"

	"github.com/reactivex/rxgo/v2"
)

var _ OrderlyMQ = (*FakeOrderlyMQ)(nil)

type FakeOrderlyMQ struct {
	events   chan rxgo.Item
	eventsRX rxgo.Observable

	subscribers map[string]Subscriber
}

func (f *FakeOrderlyMQ) Run() error { return nil }

func (f *FakeOrderlyMQ) Register(subscribers ...Subscriber) {
	for _, sub := range subscribers {
		f.subscribers[sub.Label()] = sub
	}
}

func (f *FakeOrderlyMQ) Push(event *Event) error {
	f.events <- rxgo.Item{V: &PipeEvent{
		event:              event,
		affectedSubscriber: f.subscribers["test"],
		observable:         nil,
		retry:              0,
	},
	}
	return nil
}

func (f *FakeOrderlyMQ) Pipe() rxgo.Observable {
	return f.eventsRX
}

func TestExchange_Start(t *testing.T) {
	sub := &TestSubscriber{handleMustOK: true}
	mq := &FakeOrderlyMQ{
		events:      make(chan rxgo.Item, 128),
		subscribers: map[string]Subscriber{},
	}
	mq.eventsRX = rxgo.FromChannel(mq.events)

	ex := New(context.TODO(), mq)
	ex.AddSubscriber(sub)
	ex.Start()

	go func() {
		ex.Push(&Event{
			Action: ActionCreate,
			Context: &thirdparty.Context{
				Label:     "wechat",
				Namespace: "xxxx",
			},
			Offset:      0,
			Users:       nil,
			Departments: nil,
		})
	}()

	time.Sleep(1 * time.Second)
	assert.True(t, sub.handleExecuted)
}
