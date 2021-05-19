package db

import (
	"database/sql"

	"github.com/luanruisong/borm/sqlbuilder"
)

type (
	dataBase struct {
		db *sql.DB
	}
)

func (d *dataBase) SetMaxOpenConns(n int) {
	d.db.SetMaxOpenConns(n)
}

func (d *dataBase) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(sql, args...)
}

func (d *dataBase) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(sql, args...)
}

func (d *dataBase) QueryRow(sql string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(sql, args...)
}

func (d *dataBase) Select(s ...string) Selector {
	return NewSelector(d.db, sqlbuilder.Select(s...))
}

func (d *dataBase) From(s string) Selector {
	return NewSelector(d.db, sqlbuilder.SelectFrom(s))
}

func (d *dataBase) DeleteFrom(tableName string) Deleter {
	return NewDeleter(d.db, sqlbuilder.DeleteFrom(tableName))
}

func (d *dataBase) AutoDelete(i interface{}) (sql.Result, error) {
	sb, err := sqlbuilder.AutoDelete(i)
	if err != nil {
		return nil, err
	}
	return NewDeleter(d.db, sb).Exec()
}

func (d *dataBase) UpdateFrom(tableName string) Updater {
	return NewUpdate(d.db, sqlbuilder.UpdateFrom(tableName))
}

func (d *dataBase) AutoUpdate(i interface{}) (sql.Result, error) {
	sb, err := sqlbuilder.AutoUpdate(i)
	if err != nil {
		return nil, err
	}
	return NewUpdate(d.db, sb).Exec()
}

func (d *dataBase) InsertInto(t string) Inserter {
	return NewInserter(d.db, sqlbuilder.InsertInto(t))
}

func (d *dataBase) AutoInsert(t interface{}) (sql.Result, error) {
	return NewInserter(d.db, sqlbuilder.AutoInsert(t)).Exec()
}
