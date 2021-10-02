package user

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/molizz/farm/model/db"
)

func New(queryer sqlx.Ext) *Queryer {
	return &Queryer{
		Model: db.NewModel(queryer, tableName),
	}
}

func (q *Queryer) GetByID(queryer sqlx.Ext, id int64) (*FarmUser, error) {
	where := squirrel.Eq{
		"id": id,
	}
	users := make([]*FarmUser, 0, 1)
	err := q.Select(columns, where, users)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		return users[0], nil
	}
	return nil, nil
}

func (q *Queryer) ListByThirdPartyID(queryer sqlx.Ext, thirdpartyID int64) ([]*FarmUser, error) {
	where := squirrel.Eq{
		"third_party_id": thirdpartyID,
	}
	users := make([]*FarmUser, 0)
	err := q.Select(columns, where, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
