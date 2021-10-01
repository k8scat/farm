package db

import (
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

func (m *Model) Select(dest interface{}, columns []string, sqlizer sq.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return errors.WithStack(err)
	}
	err = sqlx.Select(m.queryer, dest, query, args...)
	return errors.WithStack(err)
}
