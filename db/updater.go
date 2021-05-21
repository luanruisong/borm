package db

import (
	"database/sql"

	"github.com/luanruisong/borm/sqlbuilder"
)

type (
	updater struct {
		sb   sqlbuilder.UpdateBuilder
		exec SqlExecutor
	}
)

func (i *updater) Set(key string, value interface{}) Updater {
	i.sb.Set(key, value)
	return i
}

func (i *updater) Or(sql string, value ...interface{}) Updater {
	i.sb.Or(sql, value...)
	return i
}

func (i *updater) Where(sql string, value ...interface{}) Updater {
	i.sb.Where(sql, value...)
	return i
}

func (i *updater) And(sql string, value ...interface{}) Updater {
	i.sb.And(sql, value...)
	return i
}

func (i *updater) Exec() (sql.Result, error) {
	sb := i.sb
	return i.exec.Exec(sb.Sql(), sb.Args()...)
}

func NewUpdate(exec SqlExecutor, sb sqlbuilder.UpdateBuilder) Updater {
	return &updater{
		sb:   sb,
		exec: exec,
	}
}
