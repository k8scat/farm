package thirdparty

type Config struct {
	Values map[string]interface{} // key/value pairs
}

type Configer interface {
	// 初始化配置信息
	// 对于一些三方，可能会有一些初始化需要的配置信息
	// 例如企业微信的通讯录中的消息通知token、secret，生成一次即可
	InitConfig() error
	// 获取配置信息
	GetConfig() (*Config, error)
	// 设置配置信息
	SetConfig(*Config) error
	// 当有修改配置时，响应该方法
	OnChangeConfig(func(*Config)) error
}
