package sqlbuilder

import "fmt"

type (
	deleteBuilder struct {
		tableName string
		whereSql  string
		whereArgs []interface{}
	}
)

func (d *deleteBuilder) Select(s ...string) SqlBuilder {
	panic("implement me")
}

func (d *deleteBuilder) Sql() string {
	return fmt.Sprintf("delete from %s where %s", d.tableName, d.whereSql)
}

func (d *deleteBuilder) Args() []interface{} {
	return d.whereArgs
}

func (d *deleteBuilder) Set(key string, value interface{}) SqlBuilder {
	panic("implement me")
}

func (d *deleteBuilder) From(tableName string) SqlBuilder {
	d.tableName = tableName
	return d
}

func (is *deleteBuilder) Where(sql string, value ...interface{}) SqlBuilder {
	is.whereSql = sql
	is.whereArgs = value
	return is
}

func DeleteFrom(tableName string) SqlBuilder {
	return new(deleteBuilder).From(tableName)
}
