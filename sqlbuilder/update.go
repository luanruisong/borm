package sqlbuilder

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/luanruisong/borm/reflectx"
)

type (
	updateBuilder struct {
		sync.Once
		tableName string
		set       map[string]interface{}
		setValues string
		args      []interface{}
		where     whereBuilder
	}
)

func (w *updateBuilder) And(sql string, value ...interface{}) UpdateBuilder {
	w.where.And(sql, value)
	return w
}

func (w *updateBuilder) Or(sql string, value ...interface{}) UpdateBuilder {
	w.where.Or(sql, value)
	return w
}

func (w *updateBuilder) Where(sql string, value ...interface{}) UpdateBuilder {
	w.where.Where(sql, value)
	return w
}

func (is *updateBuilder) Set(k string, v interface{}) UpdateBuilder {
	is.Do(func() {
		is.set = make(map[string]interface{})
	})
	is.set[k] = v
	return is
}

func (is *updateBuilder) From(tableName string) UpdateBuilder {
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
	sql := fmt.Sprintf("update from %s set %s", is.tableName, is.setValues)
	if !is.where.Empty() {
		sql += fmt.Sprintf(" where %s", is.where.Sql())
	}
	return sql
}

func (is *updateBuilder) Args() []interface{} {
	w := is.where.Args()
	return append(is.args, w...)
}

func UpdateFrom(tableName string) UpdateBuilder {
	return new(updateBuilder).From(tableName)
}

func StructToWhere(i interface{}) (where string, whereArgs []interface{}) {
	var (
		whereList []string
	)
	_ = reflectx.StructRange(i, func(t reflect.StructField, v reflect.Value) error {
		//获取tag，也就是自定义的column
		column := ColumnName(t)
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

func HalfAutoUpdate(set interface{}, where interface{}) (UpdateBuilder, error) {
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
		sb.Where(whereSql, whereArgs...)
	}
	return sb, nil
}

func AutoUpdate(i interface{}) (UpdateBuilder, error) {
	//老样子 先拿表名来生成一个sqlbuilder
	tName := TableName(i)
	if len(tName) == 0 {
		return nil, errors.New("can not find table name")
	}
	sb := UpdateFrom(tName)

	ic := 0 //有效列计数
	_ = reflectx.StructRange(i, func(t reflect.StructField, v reflect.Value) error {
		if ws := whereSql(t); len(ws) > 0 {
			sb.And(ws, v.Interface())
		} else {
			//获取tag，也就是自定义的column
			column := ColumnName(t)
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
