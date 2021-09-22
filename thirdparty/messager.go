package thirdparty

// 消息发送者
type MessageSender interface {
	// 是否开启消息发送
	Enable() bool
	// 发送消息
	Send(receviers []string, body string) error
}

// 消息推送
type Messager interface {
	// 是否有发送消息的能力
	// eg. LDAP是没有发送推送能力的返回false
	// eg. 企业微信是有消息推送能力的返回true
	HasSender() bool
	// 获取消息发送者，对消息进行推送
	GetSender() MessageSender
}
