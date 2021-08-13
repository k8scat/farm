package thridparty

type ThridPartyUser struct {
	Primary    string                 // 主键
	Attributes map[string]interface{} // 用户属性
}

type ThridParty interface {
	Name() string
	PullUsers()
	PUllDepts()
}

type Manager interface {
	AddThirdParty(ThridParty)
}
