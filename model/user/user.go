package user

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/molizz/farm/model/db"
)

func New(queryer sqlx.Ext) *Queryer {
	return &Queryer{
		Model: db.NewModel(queryer),
	}
}

func (q *Queryer) GetByID(queryer sqlx.Ext, id int64) (*User,error) {
	squirrel.Select(columns...)
	q.Select(dest , sqlizer squirrel.Sqlizer)
	return nil
}
