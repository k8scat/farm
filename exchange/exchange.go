package exchange

import "github.com/molizz/farm/thirdparty"

// Subscriber 同步后的数据进行整理后拆成单独的event，并推送到订阅者
type Subscriber interface {
}

// Exchange
// 订阅者管理中心
// 将消息推送到订阅者
// 根据ack策略进行处理，保持有序性 & 最终一致性
type Exchange struct {
	// filter 这里的filter将处理消息的筛选，被推送的消息将从这里得到过滤
	filter thirdparty.ThirdPartyUserFilter
}

func (e *Exchange) AddSubscriber(sub Subscriber) {

}
