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

func (m *Model) Insert(columns []string, values []interface{}) error {
	query, args, err := sq.Insert(m.tableName).Columns(columns...).Values(values...).ToSql()
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = m.queryer.Exec(query, args...)
	return errors.WithStack(err)
}

func (m *Model) Select(columns []string, where sq.Sqlizer, dest interface{}) error {
	query, args, err := sq.Select(columns...).
		From(m.tableName).
		Where(where).ToSql()
	if err != nil {
		return errors.WithStack(err)
	}

	err = sqlx.Select(m.queryer, dest, query, args...)
	return errors.WithStack(err)
}

func (m *Model) SelectOnStream(
	columns []string,
	where sq.Sqlizer,
	makeFunc func() (dest interface{}),
	callbackFunc func(dest interface{}) error) error {

	query, args, err := sq.Select(columns...).
		From(m.tableName).
		Where(where).ToSql()
	if err != nil {
		return errors.WithStack(err)
	}

	rows, err := m.queryer.Queryx(query, args...)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		dest := makeFunc()
		err = rows.StructScan(dest)
		if err != nil {
			return errors.WithStack(err)
		}

		err = callbackFunc(dest)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (m *Model) Update(values map[string]interface{}, where sq.Sqlizer) (affected int64, err error) {
	query, args, err := sq.Update(m.tableName).SetMap(values).Where(where).ToSql()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	result, err := m.queryer.Exec(query, args...)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	affected, _ = result.RowsAffected()
	return
}
