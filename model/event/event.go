package event

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/molizz/farm/model/db"
	"github.com/pkg/errors"
)

func New(queryer sqlx.Ext) *Queryer {
	return &Queryer{
		Model: db.NewModel(queryer, tableName),
	}
}

func (q *Queryer) GetByID(id int64) (*Event, error) {
	var events = make([]*Event, 0)
	var where = squirrel.Eq{"id": id}
	err := q.Select(columns, where, &events)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(events) == 0 {
		return events[0], nil
	}
	return nil, nil
}

func (q *Queryer) ListByNamespaceOnStream(namespace string, minID int64, fn func(*Event)) error {
	var events = make([]*Event, 0)
	var where = squirrel.And{
		squirrel.Eq{"namespace": namespace},
		squirrel.Gt{"id": minID},
	}

	err := q.Select(columns, where, &events)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return events, nil
}
