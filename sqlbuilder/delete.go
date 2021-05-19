package sqlbuilder

import (
	"errors"
	"fmt"
)

type (
	deleteBuilder struct {
		tableName string
		where     whereBuilder
	}
)

func (w *deleteBuilder) And(sql string, value ...interface{}) DeleteBuilder {
	w.where.And(sql, value)
	return w
}

func (w *deleteBuilder) Or(sql string, value ...interface{}) DeleteBuilder {
	w.where.Or(sql, value)
	return w
}

func (w *deleteBuilder) Where(sql string, value ...interface{}) DeleteBuilder {
	w.where.Where(sql, value)
	return w
}

func (d *deleteBuilder) Sql() string {
	sql := fmt.Sprintf("delete from %s", d.tableName)
	if !d.where.Empty() {
		sql += fmt.Sprintf(" where %s", d.where.Sql())
	}
	return sql
}

func (d *deleteBuilder) Args() []interface{} {
	return d.where.Args()
}

func (d *deleteBuilder) From(tableName string) DeleteBuilder {
	d.tableName = tableName
	return d
}

func DeleteFrom(tableName string) DeleteBuilder {
	return new(deleteBuilder).From(tableName)
}

func AutoDelete(i interface{}) (DeleteBuilder, error) {
	//老样子 先拿表名来生成一个sqlbuilder
	tName := TableName(i)
	if len(tName) == 0 {
		return nil, errors.New("can not find table name")
	}
	sb := DeleteFrom(tName)
	where := AutoWhere(i)
	sb.Where(where.Sql(), where.Args()...)
	return sb, nil
}
