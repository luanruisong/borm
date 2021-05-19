package db

import (
	"database/sql"

	"github.com/luanruisong/borm/sqlbuilder"
)

type (
	inserter struct {
		sb sqlbuilder.InsertBuilder
		db *sql.DB
	}
)

func (i *inserter) Exec() (sql.Result, error) {
	sb := i.sb
	return i.db.Exec(sb.Sql(), sb.Args()...)
}

func (i *inserter) Values(i2 interface{}) Inserter {
	i.sb.Values(i2)
	return i
}

func NewInserter(db *sql.DB, sb sqlbuilder.InsertBuilder) Inserter {
	return &inserter{
		sb: sb,
		db: db,
	}
}
