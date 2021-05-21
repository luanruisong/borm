package db

import (
	"database/sql"

	"github.com/luanruisong/borm/sqlbuilder"
)

type (
	inserter struct {
		sb   sqlbuilder.InsertBuilder
		exec SqlExecutor
	}
)

func (i *inserter) Exec() (sql.Result, error) {
	sb := i.sb
	return i.exec.Exec(sb.Sql(), sb.Args()...)
}

func (i *inserter) Values(i2 interface{}) Inserter {
	i.sb.Values(i2)
	return i
}

func NewInserter(exec SqlExecutor, sb sqlbuilder.InsertBuilder) Inserter {
	return &inserter{
		sb:   sb,
		exec: exec,
	}
}
