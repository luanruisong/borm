package db

import (
	"database/sql"

	"github.com/luanruisong/borm/sqlbuilder"
)

type (
	deleter struct {
		sb   sqlbuilder.DeleteBuilder
		exec SqlExecutor
	}
)

func (i *deleter) Or(sql string, value ...interface{}) Deleter {
	i.sb.Or(sql, value...)
	return i
}

func (i *deleter) Where(sql string, value ...interface{}) Deleter {
	i.sb.Where(sql, value...)
	return i
}

func (i *deleter) And(sql string, value ...interface{}) Deleter {
	i.sb.And(sql, value...)
	return i
}

func (i *deleter) Exec() (sql.Result, error) {
	sb := i.sb
	return i.exec.Exec(sb.Sql(), sb.Args()...)
}

func NewDeleter(exec SqlExecutor, sb sqlbuilder.DeleteBuilder) Deleter {
	return &deleter{
		sb:   sb,
		exec: exec,
	}
}
