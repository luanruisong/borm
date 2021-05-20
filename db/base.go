package db

import (
	"database/sql"

	"github.com/luanruisong/borm/iterator"
)

type (
	Executor interface {
		Exec() (sql.Result, error)
	}
	Updater interface {
		Executor
		Set(key string, value interface{}) Updater
		Where(sql string, value ...interface{}) Updater
		And(sql string, value ...interface{}) Updater
		Or(sql string, value ...interface{}) Updater
	}
	Deleter interface {
		Executor
		Where(sql string, value ...interface{}) Deleter
		And(sql string, value ...interface{}) Deleter
		Or(sql string, value ...interface{}) Deleter
	}

	Inserter interface {
		Executor
		Values(interface{}) Inserter
	}
	Selector interface {
		iterator.Iterator
		Where(sql string, value ...interface{}) Selector
		And(sql string, value ...interface{}) Selector
		Or(sql string, value ...interface{}) Selector
		From(tableName string) Selector
		Select(...string) Selector
		OrderBy(string) Selector
		GroupBy(string) Selector
		Limit(int64) Selector
		Offset(int64) Selector
	}
	DB interface {
		SetMaxOpenConns(n int)
		Exec(sql string, args ...interface{}) (sql.Result, error)
		Query(sql string, args ...interface{}) (*sql.Rows, error)
		QueryRow(sql string, args ...interface{}) *sql.Row

		InsertInto(t string) Inserter
		AutoInsert(t interface{}) (sql.Result, error)
		UpdateFrom(tableName string) Updater
		AutoUpdate(i interface{}) (sql.Result, error)
		DeleteFrom(tableName string) Deleter
		AutoDelete(i interface{}) (sql.Result, error)

		Select(...string) Selector
		SelectFrom(string) Selector
	}

	Connector interface {
		DriverName() string
		ConnStr() string
	}
)

func NewDB(driver, connStr string) (DB, error) {
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	return &dataBase{db: db}, nil
}
