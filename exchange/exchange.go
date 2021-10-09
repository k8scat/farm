package exchange

// Subscriber 同步后的数据进行整理后拆成单独的event，并推送到订阅者
// 订阅者应该可能至少只有一个实例来接收farm产生的事件
type Subscriber interface {
	Label() string             // 订阅者的标签
	Actions() []Action         // 订阅者的事件类型
	Handle(event *Event) error // 处理event
	LastOffset() uint64        // 订阅者的最后一个offset
}

type OrderlyMQ interface {
	// 注册订阅者
	Subscribers(...Subscriber) error
	// 将事件推给MQ
	Push(*Event) error
}

// Exchange
// 订阅者管理中心
// 将消息推送mq（mq将根据订阅者的情况，选择性的推送到订阅者）
type Exchange struct {
	// 有序的队列
	mq OrderlyMQ
}

func New(mq OrderlyMQ) *Exchange {
	e := &Exchange{
		mq: mq,
	}
	return e
}

func (e *Exchange) AddSubscriber(sub Subscriber) {
	e.mq.Subscribers(sub)
}

func (e *Exchange) Push(event *Event) error {
	// event
	return e.mq.Push(event)
}
