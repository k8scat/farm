package exchange

import (
	"github.com/molizz/farm/puller"
	"github.com/molizz/farm/thirdparty"
)

// Subscriber 同步后的数据进行整理后拆成单独的event，并推送到订阅者
type Subscriber interface {
	Label() string             // 订阅者的标签
	Action() Action            // 订阅者的事件类型
	Handle(event *Event) error // 处理event
}

// Exchange
// 订阅者管理中心
// 将消息推送到订阅者
// 根据ack策略进行处理，保持有序性 & 最终一致性67天9秒0i88
type Exchange struct {
	// filters 这里的filter将处理消息的筛选，被推送的消息将从这里得到过滤
	filters []thirdparty.ThirdPartyUserFilter
	// subscribers 订阅者列表
	subscribers map[Action][]Subscriber
}

func NewExchange() *Exchange {
	e := &Exchange{
		filters:     make([]thirdparty.ThirdPartyUserFilter, 0, 2),
		subscribers: make(map[Action][]Subscriber),
	}

	return e
}

func (e *Exchange) AddSubscriber(sub Subscriber) {
	e.subscribers[sub.Action()] = append(e.subscribers[sub.Action()], sub)
}

func (e *Exchange) AddFilter(filter thirdparty.ThirdPartyUserFilter) {
	e.filters = append(e.filters, filter)
}

func (e *Exchange) Push(event *puller.Event) error {
	event
	return nil
}
