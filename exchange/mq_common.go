package exchange

import (
	"time"

	"github.com/reactivex/rxgo/v2"
)

const (
	RETRY_COUNT    = 3
	RETRY_INTERVAL = 3 * time.Second
)

type MQConfig struct {
	// 重试次数
	RetryCount int
	// 每次重试间隔
	RetryInterval time.Duration
}

var DefaultMQConfig = &MQConfig{
	RetryCount:    RETRY_COUNT,
	RetryInterval: RETRY_INTERVAL,
}

type PipeEvent struct {
	// event 当前要被推送给用户的Event
	event *Event
	// affectedSubscribers 被通知的Subscriber
	affectedSubscriber Subscriber
	// obs 管理工具
	obs rxgo.Observable
	// retry is a retry count
	retry int
}

func (p *PipeEvent) Wait(shouldFunc func(*Event) error) error {
	err := shouldFunc(p.event)
	if err != nil {

		p.obs.StartWith(rxgo.Just(p)())
		p.retry++
		// 用户处理Event出错
		// 将Event放入队列最前面，并延迟 DefaultMQConfig.RetryCount 秒，等待下次重试
		if p.retry > DefaultMQConfig.RetryCount {
			return nil
		}
		time.Sleep(DefaultMQConfig.RetryInterval)
		return err
	} else {
		// 重置 subscriber 的offset
		_ = p.affectedSubscriber.SetOffset(p.event.Offset)
	}

	return nil
}
