package exchange

// Subscriber 同步后的数据进行整理后拆成单独的event，并推送到订阅者
// 订阅者应该可能至少只有一个实例来接收farm产生的事件
type Subscriber interface {
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
	// 运行mq
	Run() error
	// Subscribers 注册订阅者
	Subscribers(...Subscriber)
	// Push 将事件推给MQ
	Push(*Event) error
	// Pipe 从MQ中获取事件
	Pipe() *PipeEvent
}

// Exchange 事件管理、推送中心
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

func (e *Exchange) Start() {
	go e.mq.Run()
}

func (e *Exchange) AddSubscriber(sub Subscriber) {
	e.mq.Subscribers(sub)
}

func (e *Exchange) Push(event *Event) error {
	// push event
	return e.mq.Push(event)
}
