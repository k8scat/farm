package thridparty

type ThridPartyUser struct {
	// 一旦确认该标识，则不允许被修改（如果修改了则需要重新同步）
	// 所以定义该字段要谨慎
	Primary    string                 // 主键，唯一性标识
	Name       string                 // 客户名称
	Attributes map[string]interface{} // 用户属性
}

type ThridPartyDepartment struct {
	// 一旦确认该标识，则不允许被修改，如果修改了则需要重新同步
	// 所以定义该字段要谨慎
	Primary    string                 // 主键，唯一性标识
	ParentID   string                 // 父部门id
	Name       string                 // 部门名称
	Attributes map[string]interface{} // 部门属性
}

type InjectIncrementUserCallbackFunc func(users []*ThridPartyUser, depts []*ThridPartyDepartment) error

// 拉取用户信息
type ThridPartyUserPuller interface {
	// 三方唯一性标识
	Label() string
	// 拉取用户
	PullUsers() ([]*ThridPartyUser, error)
	// 拉取部门
	PullDepts() ([]*ThridPartyDepartment, error)
	// 是否支持增量拉取
	// 对于像微信、钉钉、飞书等支持增量拉取的第三方，可以返回true
	IsIncrement() bool
	// 当有增量信息时，调用fn传递给Synchronizer
	InjectPullIncrementCallback(fn InjectIncrementUserCallbackFunc) error
}

// 三方用户过滤器
// TODO 实现默认过滤器（仅进行部门过滤）
type ThridPartyUserFilter interface {
	// 是否有过滤器（如果没有过滤器，则返回false）
	HasFilter() bool
	// 过滤规则
	Filter([]*ThridPartyUser, []*ThridPartyDepartment) ([]*ThridPartyUser, []*ThridPartyDepartment, error)
}

// 三方拉取器
type ThridPartyPuller interface {
	// 获取三方的拉取能力
	GetPuller() ThridPartyUserPuller
	// 获取三方的过滤器
	GetFilter() ThridPartyUserFilter
	// 定时任务
	// true表示需要定时任务，并返回定时任务的时间cron表达式
	// false表示不需要定时任务
	Cron() (bool, string)
}

type Synchronizer interface {
	// 注册三方同步器
	// 当有三方同步器注册时，会调用该方法
	// 如果返回错误，则表示该三方同步器重复注册
	RegisterPuller(puller ThridPartyPuller) error
	Do() error
}
