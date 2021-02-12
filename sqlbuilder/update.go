package sqlbuilder

import (
	"fmt"
	"strings"
	"sync"
)

type (
	updateBuilder struct {
		sync.Once
		tableName string
		set       map[string]interface{}
		setValues string
		whereSql  string
		whereArgs []interface{}
		args      []interface{}
	}
)

func (is *updateBuilder) Select(s ...string) SqlBuilder {
	panic("implement me")
}

func (is *updateBuilder) Where(sql string, value ...interface{}) SqlBuilder {
	is.whereSql = sql
	is.whereArgs = value
	return is
}

func (is *updateBuilder) Set(k string, v interface{}) SqlBuilder {
	is.Do(func() {
		is.set = make(map[string]interface{})
	})
	is.set[k] = v
	return is
}

func (is *updateBuilder) From(tableName string) SqlBuilder {
	is.tableName = tableName
	return is
}

func (is *updateBuilder) Sql() string {
	if len(is.setValues) == 0 {
		sets := make([]string, 0)
		for i, v := range is.set {
			sets = append(sets, fmt.Sprintf("%s = ?", i))
			is.args = append(is.args, v)
		}
		is.setValues = strings.Join(sets, ",")
	}
	return fmt.Sprintf("update from %s set %s where %s", is.tableName, is.setValues, is.whereSql)
}

func (is *updateBuilder) Args() []interface{} {
	return append(is.args, is.whereArgs...)
}

func UpdateFrom(tableName string) SqlBuilder {
	return new(updateBuilder).From(tableName)
}
