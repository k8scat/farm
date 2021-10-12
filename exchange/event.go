package exchange

import (
	"encoding/json"
	"time"

	eventModel "github.com/molizz/farm/model/event"
	"github.com/molizz/farm/thirdparty"
)

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
	Action Action `json:"action"`
	// 事件所属上下文(三方类型、三方的所属命名空间)
	Context *thirdparty.Context `json:"context"`
	// 当前event的offset
	// 只有从mq中取出来时，才会设置该变量
	Offset uint64 `json:"-"`
	// 用户数据
	Users []*thirdparty.ThirdPartyUser `json:"users"`
	// 部门数据
	Departments []*thirdparty.ThridPartyDepartment `json:"departments"`
}

func (e *Event) ToJSON() string {
	raw, _ := json.Marshal(e)
	return string(raw)
}

func (e *Event) ToModel() *eventModel.Event {
	return &eventModel.Event{
		Namespace:  e.Context.Namespace,
		Payload:    e.ToJSON(),
		CreateTime: time.Now().Unix(),
	}
}
