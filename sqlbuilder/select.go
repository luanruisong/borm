package sqlbuilder

import (
	"fmt"
	"strings"
)

type (
	selectBuilder struct {
		tableName string
		column    []string
		orderBy   []string
		groupBy   []string
		limit     int64
		offset    int64
		whereSql  string
		whereArgs []interface{}
	}
)

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
	if len(is.whereSql) > 0 {
		where = fmt.Sprintf("where %s", is.whereSql)
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
	return s.whereArgs
}

func (s *selectBuilder) Where(sql string, value ...interface{}) Selector {
	s.whereSql = sql
	s.whereArgs = value
	return s
}

func SelectFrom(tableName string) Selector {
	return &selectBuilder{
		tableName: tableName,
	}
}
