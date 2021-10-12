package exchange

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/molizz/farm/model/db"
	subscriberModel "github.com/molizz/farm/model/subscriber"
)

// Offset 默认的Offset的能力
// 它可以被实现了Subscriber的类型组合使用
type Offset struct {
	// 订阅者的标识
	subscriberLabel string
	// 偏移
	offset uint64
	// 锁offset的设置，保证offset在多个event触发时的唯一性
	offsetMu sync.Mutex
}

func NewDefaultOffset(subscriberLabel string) *Offset {
	return &Offset{
		subscriberLabel: subscriberLabel,
		offset:          0,
	}
}

// LastOffset 订阅者的最后一个offset
func (o *Offset) LastOffset() uint64 {
	o.offsetMu.Lock()
	defer o.offsetMu.Unlock()

	if o.offset > 0 {
		return o.offset
	}

	// 从数据库中取
	sub, err := subscriberModel.New(db.GetDB()).Get(o.subscriberLabel)
	if err != nil || sub == nil {
		log.Printf("Get last offset by '%s' was err: %+v", o.subscriberLabel, err)
		return 0
	}
	o.offset = sub.Offset
	return o.offset
}

// SetOffset 设置订阅者的最后一个offset
func (o *Offset) SetOffset(offset uint64) error {
	o.offsetMu.Lock()
	defer o.offsetMu.Unlock()

	dbErr := db.Transact(func(tx sqlx.Ext) error {
		affected, err := subscriberModel.New(tx).UpdateOffset(o.subscriberLabel, offset)
		if err != nil {
			return err
		}
		if affected > 0 {
			o.offset = offset
		}
		return nil
	})

	return dbErr
}
