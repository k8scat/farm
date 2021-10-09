package exchange

import "github.com/molizz/farm/thirdparty"

type (
	Action string
)

const (
	ActionCreate Action = "create"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
)

type Event struct {
	// 事件类型
	Action Action
	// 事件所属命名空间
	Context *thirdparty.Context
	// 用户数据
	Users []*thirdparty.ThirdPartyUser
	// 部门数据
	Departments []*thirdparty.ThridPartyDepartment
}
