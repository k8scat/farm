package thirdparty

// 三方需要实现的能力
// 任何三方都应该实现这些能力
type ThirdParty interface {
	// Label 三方唯一性标识
	// 例1 wechat
	// 例2 lark
	// 例3 ldap
	Label() string

	GetThirdPartyPuller() ThirdPartyPuller
	GetUserManager() UserManager
	GetMessager() Messager
	GetOAuth2() OAuth2
	GetAuthorizer() Authorizer
	GetCaller() Caller
	GetConfiger() Configer
}
