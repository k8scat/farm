package exchange

// 默认的Offset的能力
// 它可以被实现了Subscriber的类型组合使用
type Offset struct {
	// 订阅者的标识
	subscriberLabel string
	// 偏移
	offset uint64
}

func NewDefaultOffset(subscriberLabel string) *Offset {
	return &Offset{
		subscriberLabel: subscriberLabel,
		offset:          0,
	}
}

// 订阅者的最后一个offset
func (o *Offset) LastOffset() uint64 {
	// TOOD 从数据库中
	return 0
}

// 设置订阅者的最后一个offset
func (o *Offset) SetOffset(uint64) error {

	return nil
}
