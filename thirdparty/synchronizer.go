package thirdparty

// 描述信息的所属上下文
type Context struct {
	Label     string
	Namespace string
}

type ThirdPartyUser struct {
	Context *Context               // 所属上下文
	Hash    uint64                 // xxhash
	Values  map[string]interface{} // 全量的用户属性值map
}

type ThridPartyDepartment struct {
	Context  *Context               // 所属上下文
	Hash     uint64                 // xxhash
	ID       string                 // 当前ID
	ParentID string                 // 父部门ID
	Values   map[string]interface{} // 全量的部门属性值map
}

type ThirdPartyPulledPack struct {
	UsersHash uint64
	DeptsHash uint64

	Users []*ThirdPartyUser
	Depts []*ThridPartyDepartment
}

type InjectIncrementUserCallbackFunc func(puller ThirdPartyPuller, pack *ThirdPartyPulledPack) error
type InjectPullActionCallbackFunc func(puller ThirdPartyPuller) error

// ThirdPartyUserPuller 拉取用户信息
type ThirdPartyUserPuller interface {
	// Namespace 应该返回当前puller的实例级别的唯一标识
	// 例如 返回 org uuid，表示puller的拉取的用户&部门信息属于哪个组织
	// 该值也将被保存在数据库中，用于隔离不同的数据
	Namespace() string

	// UserPrimaryAttrs 拉取的用户 PrimaryAttrs 字段，该字段将作为用户的主键进行唯一性匹配
	// 必须返回 PullUsers() 中拥有的字段名，否则会出现错误
	// 这里返回数组的原因是，允许使用多个字段进行组合主键，通常返回一个即可，例如微信的openid
	UserPrimaryAttrs() []string

	// IndexAttrs 需要建立搜索能力的属性
	// 对于返回了该字段的属性，将对其进行数据库层面的快速检索能力
	// 当该字段被修改时，应重建索引（异步）
	IndexAttrs() []string

	// PullUsers 拉取用户、部门
	Pull() (*ThirdPartyPulledPack, error)

	// DepartmentPrimaryAttr 拉取的部门 Primary 字段，该字段将作为部门的主键进行唯一性匹配
	// 必须返回 PullDepts() 中拥有的字段名，否则会出现错误
	DepartmentPrimaryAttr() string

	// HasIncrement 是否支持增量拉取
	// 对于像微信、钉钉、飞书等支持增量拉取的第三方，可以返回true
	HasIncrement() bool

	// InjectPullIncrementCallback 当有增量信息时，调用fn传递给Synchronizer
	InjectPullIncrementCallback(InjectIncrementUserCallbackFunc) error

	// InjectPullActionFunc 用户触发拉取
	InjectPullActionFunc(InjectPullActionCallbackFunc) error
}

// ThirdPartyUserFilter 三方用户过滤器
// TODO 实现默认过滤器（仅进行部门过滤）
type ThirdPartyUserFilter interface {
	// HasFilter 是否有过滤器（如果没有过滤器，则返回false）
	HasFilter() bool
	// Filter 过滤规则
	Filter([]*ThirdPartyUser, []*ThridPartyDepartment) ([]*ThirdPartyUser, []*ThridPartyDepartment, error)
}

// ThirdPartyPuller 三方拉取器
type ThirdPartyPuller interface {
	// GetPuller 获取三方的拉取能力
	GetPuller() ThirdPartyUserPuller
	// GetFilter 获取三方的过滤器
	GetFilter() ThirdPartyUserFilter
	// Cron 定时任务
	// true表示需要定时任务，并返回定时任务的时间cron表达式
	// false表示不需要定时任务
	Cron() (bool, string)
}

type Synchronizer interface {
	// RegisterPuller 注册三方同步器
	// 当有三方同步器注册时，会调用该方法
	// 如果返回错误，则表示该三方同步器重复注册
	RegisterPuller(label string, puller ThirdPartyPuller) error
	PullerCount() int
	Do() error
}
