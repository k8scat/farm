package event

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/molizz/farm/model/db"
	"github.com/pkg/errors"
)

func New(queryer sqlx.Ext) *Queryer {
	return &Queryer{
		Model: db.NewModel(queryer, tableName),
	}
}

func (q *Queryer) Create(event *Event) error {
	err := q.Insert(columns[1:], []interface{}{
		event.Namespace,
		event.Payload,
		event.CreateTime,
	})

	return errors.WithStack(err)
}

func (q *Queryer) GetByID(id int64) (*Event, error) {
	var (
		events = make([]*Event, 0)
		where  = sq.Eq{"id": id}
	)

	err := q.Select(columns, where, &events)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(events) == 0 {
		return events[0], nil
	}
	return nil, nil
}

func (q *Queryer) ListByNamespaceOnStream(namespace string, minID uint64, fn func(*Event) error) error {
	var where = sq.And{
		sq.Eq{"namespace": namespace},
		sq.Gt{"id": minID},
	}

	makeFunc := func() interface{} {
		return new(Event)
	}

	callbackFunc := func(dest interface{}) error {
		e, ok := dest.(*Event)
		if ok {
			err := fn(e)
			if err != nil {
				return err
			}
		}
		return nil
	}

	err := q.SelectOnStream(columns, where, makeFunc, callbackFunc)
	return errors.WithStack(err)
}
