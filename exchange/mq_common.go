package exchange

type PipeEvent struct {
	// event 当前要被推送给用户的Event
	event *Event
	// affectedSubscribers 被通知的Subscriber
	affectedSubscribers []Subscriber
}

func (p *PipeEvent) Wait(waitFunc func(*Event) error) error {
	err := waitFunc(p.event)
	if err != nil {
		// 用户处理Event出错
		// TODO 将Event放入队列最前面，并延迟3秒，等待下次重试

		// rxgo.Retry(RETRY_COUNT, RETRY_INTERVAL, func() error {
		// }).Run()

		return err
	} else {
		// 重置 subscriber 的offset
		for _, sub := range p.affectedSubscribers {
			sub.SetOffset(p.event.Offset)
		}
	}

	return nil
}
