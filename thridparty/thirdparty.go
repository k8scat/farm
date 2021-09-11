package thridparty

// 三方需要实现的能力
// 任何三方都应该实现这些能力
type ThirdParty interface {
	ThridPartyPuller
	UserManager
}
