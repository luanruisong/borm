package sqlbuilder

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/luanruisong/borm/reflectx"
	"github.com/luanruisong/borm/stringx"
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

//自动插入结构体
func AutoInsert(i interface{}) (SqlBuilder, error) {
	//处理tableName 如果实现了table接口，按照接口
	var tName string
	if t, ok := i.(Table); ok {
		tName = t.TableName()
	} else {
		//未实现接口，取struct名称的蛇形
		tName = stringx.SnakeName(reflectx.StructName(i))
	}
	//如果是匿名struct，无table 返回找不到tableName
	if len(tName) == 0 {
		return nil, errors.New("can not find table name")
	}
	//初始化一个sqlBuilder
	sb := InsertInto(tName)
	//定义一个计数器
	ic := 0
	//使用reflectx的range函数，对参结构体进行遍历
	//错误可直接忽略（因为没有产生错误的地方）
	_ = reflectx.StructRange(i, func(t reflect.StructField, v reflect.Value) error {
		//获取tag，也就是自定义的column
		column := t.Tag.Get("db")
		if len(column) == 0 {
			//入无自定义column，取field名称的蛇形
			column = stringx.SnakeName(t.Name)
		}
		// "-" 表示忽略，空数据 也直接跳过
		if column == "-" || reflectx.IsNull(v) {
			return nil
		}
		// 书接上回，进行sql构建（单一属性）
		sb.Set(column, v.Interface())
		ic++
		return nil
	})
	//计数器为0表示无可用列构建，返回错误
	if ic == 0 {
		return nil, errors.New("no column field to sql")
	}
	return sb, nil
}
