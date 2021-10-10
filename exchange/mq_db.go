package exchange

// MQDB 基于数据库的MQ
// 如果对于持久性要求不高的话，可以考虑使用 Redis Stream 来实现 MQ
//
type MQDB struct {
}

func NewMQDB() *MQDB {
	mq := new(MQDB)
	return mq
}

func (d *MQDB) Push(e *Event) error {
	return nil
}

func (d *MQDB) Subscribers(subs ...Subscriber) error {
	return nil
}
