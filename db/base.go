package db

import (
	"database/sql"

	"github.com/luanruisong/borm/iterator"
)

type (
	BormExecutor interface {
		Exec() (sql.Result, error)
	}
	Updater interface {
		BormExecutor
		Set(key string, value interface{}) Updater
		Where(sql string, value ...interface{}) Updater
		And(sql string, value ...interface{}) Updater
		Or(sql string, value ...interface{}) Updater
	}
	Deleter interface {
		BormExecutor
		Where(sql string, value ...interface{}) Deleter
		And(sql string, value ...interface{}) Deleter
		Or(sql string, value ...interface{}) Deleter
	}

	Inserter interface {
		BormExecutor
		Values(interface{}) Inserter
	}
	Selector interface {
		iterator.Iterator
		Where(sql string, value ...interface{}) Selector
		AutoWhere(i interface{}) Selector
		And(sql string, value ...interface{}) Selector
		Or(sql string, value ...interface{}) Selector
		From(tableName string) Selector
		Select(...string) Selector
		OrderBy(string) Selector
		GroupBy(string) Selector
		Limit(int64) Selector
		Offset(int64) Selector
	}

	SqlExecutor interface {
		Exec(sql string, args ...interface{}) (sql.Result, error)
		Query(sql string, args ...interface{}) (*sql.Rows, error)
		QueryRow(sql string, args ...interface{}) *sql.Row
	}
	Executor interface {
		SqlExecutor
		InsertInto(t string) Inserter
		AutoInsert(t interface{}) (sql.Result, error)
		UpdateFrom(tableName string) Updater
		AutoUpdate(i interface{}) (sql.Result, error)
		DeleteFrom(tableName string) Deleter
		AutoDelete(i interface{}) (sql.Result, error)

		Select(...string) Selector
		SelectFrom(string) Selector

		Count(i interface{}) (int64, error)
	}

	DataBase interface {
		Executor
		SetMaxOpenConns(n int)
		Close() error

		Begin() (Tx, error)
	}
	Tx interface {
		Executor
		Comment() error
		RollBack() error
	}

	Connector interface {
		DriverName() string
		ConnStr() string
		GetPoolSize() int
	}
)

func NewDB(driver, connStr string) (DataBase, error) {
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &dataBase{
		executor: newExec(db),
		db:       db,
	}, nil
}
