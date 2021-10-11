package event

import (
	"github.com/molizz/farm/model/db"
)

var (
	tableName = "farm_event"
	columns   = []string{"id", "namespace", "payload", "create_time"}
)

type Queryer struct {
	db.Model
}

type Event struct {
	ID         int64  `db:"id"`
	Namespace  string `db:"namespace"`
	Payload    string `db:"payload"`
	CreateTime int64  `db:"create_time"`
}
