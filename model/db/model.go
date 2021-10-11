package db

import (
	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Model struct {
	queryer   sqlx.Ext
	tableName string
}

func NewModel(queryer sqlx.Ext, tableName string) Model {
	return Model{
		queryer:   queryer,
		tableName: tableName,
	}
}

func (m *Model) Select(columns []string, where sq.Sqlizer, dest interface{}) error {
	query, args, err := squirrel.Select(columns...).
		From(m.tableName).
		Where(where).ToSql()
	if err != nil {
		return errors.WithStack(err)
	}

	err = sqlx.Select(m.queryer, dest, query, args...)
	return errors.WithStack(err)
}

func (m *Model) StreamSelect(columns []string, where sq.Sqlizer, dest interface{}) error {
	query, args, err := squirrel.Select(columns...).
		From(m.tableName).
		Where(where).ToSql()
	if err != nil {
		return errors.WithStack(err)
	}

	rows, err := m.queryer.Queryx(query, args...)
	if err != nil {
		return errors.WithStack(err)
	}
	rows.

	return errors.WithStack(err)

}
