package db

import (
	"database/sql"

	"github.com/luanruisong/borm/sqlbuilder"
)

type (
	updater struct {
		sb sqlbuilder.UpdateBuilder
		db *sql.DB
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
	return i.db.Exec(sb.Sql(), sb.Args()...)
}

func NewUpdate(db *sql.DB, sb sqlbuilder.UpdateBuilder) Updater {
	return &updater{
		sb: sb,
		db: db,
	}
}
