package thridparty

// Authorizer 通过账号密码换取账号信息
type Authorizer interface {
	// 是否拥有通过账号密码访问的能力
	HasAuthorizer() bool
	// 获取账号信息
	GetAccount(account, password string) (*ThridPartyUser, error)
}
