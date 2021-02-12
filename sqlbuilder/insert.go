package sqlbuilder

import (
	"fmt"
	"strings"
	"sync"
)

type (
	insertBuilder struct {
		sync.Once
		tableName string
		set       map[string]interface{}
		setValues string
		args      []interface{}
	}
)

func (is *insertBuilder) Select(s ...string) SqlBuilder {
	panic("implement me")
}

func (is *insertBuilder) Where(sql string, value ...interface{}) SqlBuilder {
	panic("implement me")
}

func (is *insertBuilder) Set(k string, v interface{}) SqlBuilder {
	is.Do(func() {
		is.set = make(map[string]interface{})
	})
	is.set[k] = v
	return is
}

func (is *insertBuilder) From(tableName string) SqlBuilder {
	is.tableName = tableName
	return is
}

func (is *insertBuilder) Sql() string {
	if len(is.setValues) == 0 {
		sets := make([]string, 0)
		for i, v := range is.set {
			sets = append(sets, fmt.Sprintf("%s = ?", i))
			is.args = append(is.args, v)
		}
		is.setValues = strings.Join(sets, ",")
	}
	return fmt.Sprintf("insert into %s set %s", is.tableName, is.setValues)
}

func (is *insertBuilder) Args() []interface{} {
	return is.args
}

func InsertInto(tableName string) SqlBuilder {
	return new(insertBuilder).From(tableName)
}
