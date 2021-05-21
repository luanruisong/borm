package db

import (
	"database/sql"
)

type (
	tx struct {
		*executor
		db *sql.Tx
	}
)

func (d *tx) Comment() error {
	return d.db.Commit()
}

func (d *tx) RollBack() error {
	return d.db.Rollback()
}

func newTx(db *sql.Tx) Tx {
	return &tx{
		executor: newExec(db),
		db:       db,
	}
}
