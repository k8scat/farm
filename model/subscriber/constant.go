package subscriber

import (
	"github.com/molizz/farm/model/db"
)

var (
	tableName = "farm_subscriber"
	columns   = []string{"label", "offset", "update_time"}
)

type Queryer struct {
	db.Model
}

type Subscriber struct {
	Label      string `db:"label"`
	Offset     uint64 `db:"offset"`
	UpdateTime int64  `db:"update_time"`
}
