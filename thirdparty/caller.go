package thirdparty

type CallerEvent interface {
	Key() string
	Func() func(callbackBody string) error
}

// 统一托管所有的三方的消息推送
// eg. 企业微信的通信录变更后到推送
type Caller interface {
	// callback 的唯一标识
	// eg. 微信
	Label() string
	// 需要注册的回调
	// eg. 微信的通信录变更后，会产生 org_user 变更事件，此时三方如果有注册 org_user 事件，那么将会回调 CallerEvent.Func() 方法
	Callbacks() []CallerEvent
}
