package thirdparty

import (
	"strings"
)

type ThirdPartyUser struct {
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

type InjectIncrementUserCallbackFunc func(puller ThirdPartyPuller, users []*ThirdPartyUser, depts []*ThridPartyDepartment) error
type InjectPullActionCallbackFunc func(puller ThirdPartyPuller) error

// ThirdPartyUserPuller 拉取用户信息
type ThirdPartyUserPuller interface {
	// Namespace 用于区分不同的Puller
	// 例 不同的用户，都在使用企业微信时，他们的数据应该是隔离的
	Namespace() Namespace
	// UserPrimaryAttrs 拉取的用户 PrimaryAttrs 字段，该字段将作为用户的主键进行唯一性匹配
	// 必须返回 PullUsers() 中拥有的字段名，否则会出现错误
	// 这里返回数组的原因是，允许使用多个字段进行组合主键，通常返回一个即可，例如微信的openid
	UserPrimaryAttrs() []string
	// PullUsers 拉取用户
	PullUsers() ([]*ThirdPartyUser, error)

	// DepartmentPrimaryAttr 拉取的部门 Primary 字段，该字段将作为部门的主键进行唯一性匹配
	// 必须返回 PullDepts() 中拥有的字段名，否则会出现错误
	DepartmentPrimaryAttr() string
	// PullDepts 拉取部门
	PullDepts() ([]*ThridPartyDepartment, error)

	// HasIncrement 是否支持增量拉取
	// 对于像微信、钉钉、飞书等支持增量拉取的第三方，可以返回true
	HasIncrement() bool
	// InjectPullIncrementCallback 当有增量信息时，调用fn传递给Synchronizer
	InjectPullIncrementCallback(fn InjectIncrementUserCallbackFunc) error

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
	// Label 三方唯一性标识
	Label() string
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
	RegisterPuller(puller ThirdPartyPuller) error
	Do() error
}

type Namespace string

func NewNamespace(parts ...string) Namespace {
	if len(parts) <= 1 {
		panic("invalid namespace parts, must gt 1")
	}
	return Namespace(strings.Join(parts, ":"))
}
