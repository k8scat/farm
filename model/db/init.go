package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var db *sqlx.DB

func MustInit() {
	if db == nil {
		panic("farm db must init")
	}
}

func SetDB(stdDB *sql.DB, driverName string) {
	db = sqlx.NewDb(stdDB, driverName)
}

func GetDB() *sqlx.DB {
	if db == nil {
		panic("db not initialized")
	}
	return db
}

func Transact(txFunc func(tx sqlx.Ext) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if e := recover(); e != nil {
			err = errors.WithStack(tx.Rollback())
			return
		}
		if err != nil {
			err = tx.Rollback()
			return
		}

		err = tx.Commit()
		return
	}()

	err = txFunc(tx)
	return errors.WithStack(err)
}
