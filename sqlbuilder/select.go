package sqlbuilder

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/luanruisong/borm/reflectx"
)

type (
	selectBuilder struct {
		tableName string
		column    []string
		orderBy   []string
		groupBy   []string
		limit     int64
		offset    int64
		where     whereBuilder
	}
)

func (w *selectBuilder) And(sql string, value ...interface{}) Selector {
	w.where.And(sql, value)
	return w
}

func (w *selectBuilder) Or(sql string, value ...interface{}) Selector {
	w.where.Or(sql, value)
	return w
}

func (w *selectBuilder) Where(sql string, value ...interface{}) Selector {
	w.where.Where(sql, value)
	return w
}

func (is *selectBuilder) Select(s ...string) Selector {
	is.column = append(is.column, s...)
	return is
}

func (is *selectBuilder) OrderBy(s string) Selector {
	is.orderBy = append(is.orderBy, s)
	return is
}

func (is *selectBuilder) GroupBy(s string) Selector {
	is.groupBy = append(is.groupBy, s)
	return is
}

func (is *selectBuilder) Limit(i int64) Selector {
	is.limit = i
	return is
}

func (is *selectBuilder) Offset(i int64) Selector {
	is.offset = i
	return is
}

func (is *selectBuilder) From(tableName string) Selector {
	is.tableName = tableName
	return is
}

func (is *selectBuilder) Sql() string {
	var (
		column  = "*"
		where   string
		groupBy string
		orderBy string
		limit   string
		offset  string
	)
	if len(is.column) > 0 {
		column = strings.Join(is.column, ",")
	}
	if !is.where.Empty() {
		where = fmt.Sprintf("where %s", is.where.Sql())
	}
	if len(is.groupBy) > 0 {
		groupBy = fmt.Sprintf("group by %s", strings.Join(is.groupBy, ","))
	}
	if len(is.orderBy) > 0 {
		orderBy = fmt.Sprintf("order by %s", strings.Join(is.orderBy, ","))
	}
	if is.limit > 0 {
		limit = fmt.Sprintf("limit %d", is.limit)
	}
	if is.offset > 0 {
		offset = fmt.Sprintf("offset %d", is.offset)
	}
	return fmt.Sprintf("select %s from %s %s %s %s %s %s", column, is.tableName, where, groupBy, orderBy, limit, offset)
}

func (s *selectBuilder) Args() []interface{} {
	return s.where.Args()
}

func SelectFrom(tableName string) Selector {
	return &selectBuilder{
		tableName: tableName,
	}
}

func AutoWhere(i interface{}) (where string, whereArgs []interface{}) {
	var (
		whereList []string
	)
	_ = reflectx.StructRange(i, func(t reflect.StructField, v reflect.Value) error {
		if sql := whereSql(t); len(sql) > 0 {
			whereList = append(whereList, sql)
			whereArgs = append(whereArgs, v.Interface())
		}
		return nil
	})
	where = strings.Join(whereList, " and ")
	return
}

func whereSql(t reflect.StructField) string {
	where := t.Tag.Get(DbwTag)
	if len(where) == 0 {
		return ""
	}
	column := ColumnName(t)
	return fmt.Sprintf("%s %s ?", column, whereFlag(where))
}

func whereFlag(flag string) string {
	switch flag {
	case DbwEqual:
		return "="
	case DbwGreaterThan:
		return ">"
	case DbwLessThan:
		return "<"
	default:
		panic("where flag " + flag + " not support")
	}
}

func AutoSelect(i interface{}) (Selector, error) {
	tName := TableName(i)
	if len(tName) == 0 {
		return nil, errors.New("can not find table name")
	}
	sb := SelectFrom(tName)

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
			sb.Select(column)
			ic++
		}
		return nil
	})
	//计数器为0表示无可用列构建，处理 *查询
	if ic == 0 {
		sb.Select("*")
	}
	return sb, nil
}
