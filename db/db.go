package db

import (
	"database/sql"
)

type (
	dataBase struct {
		*executor
		db *sql.DB
	}
)

func (d *dataBase) Close() error {
	return d.db.Close()
}

func (d *dataBase) SetMaxOpenConns(n int) {
	d.db.SetMaxOpenConns(n)
}

func (d *dataBase) Begin() (Tx, error) {
	sqlTx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}
	return newTx(sqlTx), nil
}
