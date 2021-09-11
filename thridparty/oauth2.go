package thridparty

type OAuth2Config struct {
	// The OAuth2 client ID.
	ClientId string
	// The OAuth2 client secret.
	ClientSecret string

	// The OAuth2 client redirect URL.
	RedirectUrl string
	// The OAuth2 client scope.
	Scopes []string
	// The OAuth2 client state.
	AuthUrl string
	// The OAuth2 client token URL.
	TokenUrl string
	// The OAuth2 client user info URL.
	ApiUrl string
}

type OAuth2Token struct {
	// The OAuth2 access token.
	AccessToken string
	// The OAuth2 refresh token.
	RefreshToken string
	// The OAuth2 token type.
	TokenType string
	// The OAuth2 token expiration time.
	ExpiryTime int64
}

type OAuth2UserInfo = ThridPartyUser

type OAuth2 interface {
	// 是否有OAuth2的能力
	HasOAuth2() bool
	// 获取OAuth2的配置
	GetOAuth2Config() *OAuth2Config
	// 获取OAuth2的token
	GetOAuth2Token() *OAuth2Token
	// 获取OAuth2的用户信息
	GetOAuth2UserInfoByToken(token *OAuth2Token) *OAuth2UserInfo
}
