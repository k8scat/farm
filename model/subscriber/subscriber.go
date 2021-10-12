package subscriber

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/molizz/farm/model/db"
)

func New(queryer sqlx.Ext) *Queryer {
	return &Queryer{
		Model: db.NewModel(queryer, tableName),
	}
}

func (q *Queryer) Get(label string) (*Subscriber, error) {
	where := sq.Eq{"label": label}
	dest := make([]*Subscriber, 0, 1)
	err := q.Select(columns, where, dest)
	if err != nil {
		return nil, err
	}
	if len(dest) > 0 {
		return dest[0], nil
	}
	return nil, nil
}

func (q *Queryer) UpdateOffset(label string, offset uint64) (affected int64, err error) {
	values := map[string]interface{}{
		"offset":      offset,
		"update_time": time.Now().Unix(),
	}
	where := sq.And{
		sq.Eq{
			"label": label,
		},
		sq.Lt{
			"offset": offset,
		},
	}
	affected, err = q.Update(values, where)
	return
}
