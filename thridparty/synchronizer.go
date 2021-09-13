package thridparty

type ThridPartyUser struct {
	// 一旦确认该标识，则不允许被修改（如果修改了则需要重新同步）
	// 所以定义该字段要谨慎
	Primaries    []string               // 主键，唯一性标识
	PrimaryValue string                 // 主键的值
	Hash         uint64                 // xxhash
	Attributes   map[string]interface{} // 用户属性
}

type ThridPartyDepartment struct {
	// 一旦确认该标识，则不允许被修改，如果修改了则需要重新同步
	// 所以定义该字段要谨慎
	Primary      string                 // 主键，唯一性标识
	PrimaryValue string                 // 主键的值
	Hash         uint64                 // xxhash
	ParentID     string                 // 父部门id
	Attributes   map[string]interface{} // 部门属性
}

type InjectIncrementUserCallbackFunc func(puller ThridPartyPuller, users []*ThridPartyUser, depts []*ThridPartyDepartment) error
type InjectPullActionCallbackFunc func(puller ThridPartyPuller) error

// 拉取用户信息
type ThridPartyUserPuller interface {
	// 拉取的用户 PrimaryAttrs 字段，该字段将作为用户的主键进行唯一性匹配
	// 必须返回 PullUsers() 中拥有的字段名，否则会出现错误
	// 这里返回数组的原因是，允许使用多个字段进行组合主键，通常返回一个即可，例如微信的openid
	UserPrimaryAttrs() []string
	// 拉取用户
	PullUsers() ([]*ThridPartyUser, error)

	// 拉取的部门 Primary 字段，该字段将作为部门的主键进行唯一性匹配
	// 必须返回 PullDepts() 中拥有的字段名，否则会出现错误
	DepartmentPrimaryAttr() string
	// 拉取部门
	PullDepts() ([]*ThridPartyDepartment, error)

	// 是否支持增量拉取
	// 对于像微信、钉钉、飞书等支持增量拉取的第三方，可以返回true
	HasIncrement() bool
	// 当有增量信息时，调用fn传递给Synchronizer
	InjectPullIncrementCallback(fn InjectIncrementUserCallbackFunc) error

	// 用户触发拉取
	InjectPullActionFunc(InjectPullActionCallbackFunc) error
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
	// 三方唯一性标识
	Label() string
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
