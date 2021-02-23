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

func StructToWhere(i interface{}) (where string, whereArgs []interface{}) {
	var (
		whereList []string
	)
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
		whereList = append(whereList, fmt.Sprintf("%s = ?", column))
		whereArgs = append(whereArgs, v.Interface())
		return nil
	})
	where = strings.Join(whereList, " and ")
	return
}

func HalfAutoUpdate(set interface{}, where interface{}) (SqlBuilder, error) {
	//老样子 先拿表名来生成一个sqlbuilder
	tName := TableName(set)
	if len(tName) == 0 {
		return nil, errors.New("can not find table name")
	}
	sb := UpdateFrom(tName)

	ic := 0 //有效列计数
	_ = reflectx.StructRange(set, func(t reflect.StructField, v reflect.Value) error {
		//获取tag，也就是自定义的column
		column := ColumnName(t)
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
	whereSql, whereArgs := StructToWhere(where)
	if len(whereSql) > 0 {
		sb.Where(whereSql, whereArgs)
	}
	return sb, nil
}

func AutoUpdate(i interface{}) (SqlBuilder, error) {
	//老样子 先拿表名来生成一个sqlbuilder
	tName := TableName(i)
	if len(tName) == 0 {
		return nil, errors.New("can not find table name")
	}
	sb := UpdateFrom(tName)

	ic := 0 //有效列计数
	_ = reflectx.StructRange(i, func(t reflect.StructField, v reflect.Value) error {
		//获取tag，也就是自定义的column
		column := ColumnName(t)
		if IsPk(t) {
			sb.Where(fmt.Sprintf("%s = ?", column), v.Interface())
		} else {
			// "-" 表示忽略，空数据 也直接跳过
			if column == "-" || reflectx.IsNull(v) {
				return nil
			}
			// 书接上回，进行sql构建（单一属性）
			sb.Set(column, v.Interface())
			ic++
		}
		return nil
	})
	//计数器为0表示无可用列构建，返回错误
	if ic == 0 {
		return nil, errors.New("no column field to sql")
	}
	return sb, nil
}
