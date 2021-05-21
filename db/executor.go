package db

import (
	"database/sql"

	"github.com/luanruisong/borm/sqlbuilder"
)

type (
	executor struct {
		exec SqlExecutor
	}
)

func (d *executor) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return d.exec.Exec(sql, args...)
}

func (d *executor) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return d.exec.Query(sql, args...)
}

func (d *executor) QueryRow(sql string, args ...interface{}) *sql.Row {
	return d.exec.QueryRow(sql, args...)
}

func (d *executor) Select(s ...string) Selector {
	return NewSelector(d.exec, sqlbuilder.Select(s...))
}

func (d *executor) SelectFrom(s string) Selector {
	return NewSelector(d.exec, sqlbuilder.SelectFrom(s))
}

func (d *executor) DeleteFrom(tableName string) Deleter {
	return NewDeleter(d.exec, sqlbuilder.DeleteFrom(tableName))
}

func (d *executor) AutoDelete(i interface{}) (sql.Result, error) {
	sb, err := sqlbuilder.AutoDelete(i)
	if err != nil {
		return nil, err
	}
	return NewDeleter(d.exec, sb).Exec()
}

func (d *executor) UpdateFrom(tableName string) Updater {
	return NewUpdate(d.exec, sqlbuilder.UpdateFrom(tableName))
}

func (d *executor) AutoUpdate(i interface{}) (sql.Result, error) {
	sb, err := sqlbuilder.AutoUpdate(i)
	if err != nil {
		return nil, err
	}
	return NewUpdate(d.exec, sb).Exec()
}

func (d *executor) InsertInto(t string) Inserter {
	return NewInserter(d.exec, sqlbuilder.InsertInto(t))
}

func (d *executor) AutoInsert(t interface{}) (sql.Result, error) {
	return NewInserter(d.exec, sqlbuilder.AutoInsert(t)).Exec()
}

func (d *executor) Count(i interface{}) (int64, error) {
	var (
		tb  = sqlbuilder.TableName(i)
		sb  = sqlbuilder.Select("count(*) as c").From(tb)
		tmp struct {
			C int64
		}
		err error
	)
	err = NewSelector(d.exec, sb).AutoWhere(i).One(&tmp)
	return tmp.C, err
}

func newExec(db SqlExecutor) *executor {
	return &executor{exec: db}
}
