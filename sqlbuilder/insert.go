package sqlbuilder

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/luanruisong/borm/reflectx"
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

func (is *insertBuilder) Set(k string, v interface{}) InsertBuilder {
	is.Do(func() {
		is.set = make(map[string]interface{})
	})
	is.set[k] = v
	return is
}

func (is *insertBuilder) From(tableName string) InsertBuilder {
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

func (is *insertBuilder) Values(i interface{}) InsertBuilder {
	//使用reflectx的range函数，对参结构体进行遍历
	//错误可直接忽略（因为没有产生错误的地方）
	_ = reflectx.StructRange(i, func(t reflect.StructField, v reflect.Value) error {
		//获取tag，也就是自定义的column
		column := reflectx.ColumnName(t)
		// "-" 表示忽略，空数据 也直接跳过
		if column == "-" || reflectx.IsNull(v) {
			return nil
		}
		// 书接上回，进行sql构建（单一属性）
		is.Set(column, v.Interface())
		return nil
	})
	return is
}

func InsertInto(tableName string) InsertBuilder {
	return new(insertBuilder).From(tableName)
}

//自动插入结构体
func AutoInsert(i interface{}) InsertBuilder {
	//处理tableName 如果实现了table接口，按照接口
	tName := TableName(i)
	//如果是匿名struct，无table 返回找不到tableName
	if len(tName) == 0 {
		return nil
	}
	//初始化一个sqlBuilder
	sb := InsertInto(tName).Values(i)
	return sb
}
