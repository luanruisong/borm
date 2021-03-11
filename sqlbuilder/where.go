package sqlbuilder

import (
	"strings"
)

type (
	whereType          uint8
	singleWhereBuilder struct {
		t    whereType
		Sql  string
		Args []interface{}
	}
	whereBuilder struct {
		w []singleWhereBuilder
	}
)

const (
	whereTypeWhere = whereType(0)
	whereTypeAnd   = whereType(1)
	whereTypeOr    = whereType(2)
)

func (wb whereBuilder) Empty() bool {
	return len(wb.w) == 0
}

func (wb *whereBuilder) And(sql string, i []interface{}) {
	wb.w = append(wb.w, singleWhereBuilder{
		whereTypeAnd,
		sql,
		i,
	})
}

func (wb *whereBuilder) Or(sql string, i []interface{}) {
	wb.w = append(wb.w, singleWhereBuilder{
		whereTypeOr,
		sql,
		i,
	})
}

func (wb *whereBuilder) Where(sql string, i []interface{}) {
	wb.w = append(wb.w, singleWhereBuilder{
		whereTypeWhere,
		sql,
		i,
	})
}

func (wb whereBuilder) Sql() string {
	ret := make([]string, 0)
	for i, v := range wb.w {
		if i != 0 {
			switch v.t {
			case whereTypeAnd:
				ret = append(ret, "and")
			case whereTypeOr:
				ret = append(ret, "or")
			}
		}
		ret = append(ret, v.Sql)
	}
	return strings.Join(ret, " ")
}

func (wb whereBuilder) Args() []interface{} {
	ret := make([]interface{}, 0)
	for _, v := range wb.w {
		ret = append(ret, v.Args...)
	}
	return ret
}
