package exchange

type Action string

const (
	ActionCreate Action = "create"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
	ActionAll    Action = "all"
)

type Event struct {
	Action Action
}
