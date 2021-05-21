package db

import (
	"github.com/luanruisong/borm/iterator"
	"github.com/luanruisong/borm/reflectx"
	"github.com/luanruisong/borm/sqlbuilder"
)

type (
	selector struct {
		sb   sqlbuilder.Selector
		exec SqlExecutor
	}
)

func (s *selector) AutoWhere(i interface{}) Selector {
	if !reflectx.IsNull(i) {
		if where := sqlbuilder.AutoWhere(i); where != nil {
			s.sb.Where(where.Sql(), where.Args()...)
		}
	}
	return s
}
func (s *selector) Where(sql string, value ...interface{}) Selector {
	s.sb.Where(sql, value...)
	return s
}

func (s *selector) And(sql string, value ...interface{}) Selector {
	s.sb.And(sql, value...)
	return s
}

func (s *selector) Or(sql string, value ...interface{}) Selector {
	s.sb.Or(sql, value...)
	return s
}

func (s *selector) From(tableName string) Selector {
	s.sb.From(tableName)
	return s
}

func (s *selector) Select(s2 ...string) Selector {
	s.sb.Select(s2...)
	return s
}

func (s *selector) OrderBy(s2 string) Selector {
	s.sb.OrderBy(s2)
	return s
}

func (s *selector) GroupBy(s2 string) Selector {
	s.sb.GroupBy(s2)
	return s
}

func (s *selector) Limit(i int64) Selector {
	s.sb.Limit(i)
	return s
}

func (s *selector) Offset(i int64) Selector {
	s.sb.Offset(i)
	return s
}

func (s *selector) All(i interface{}) error {
	rows, err := s.exec.Query(s.sb.Sql(), s.sb.Args()...)
	if err != nil {
		return err
	}
	iter := iterator.New(rows)
	return iter.All(i)
}

func (s *selector) One(i interface{}) error {
	rows, err := s.exec.Query(s.sb.Sql(), s.sb.Args()...)
	if err != nil {
		return err
	}
	iter := iterator.New(rows)
	return iter.One(i)
}

func NewSelector(exec SqlExecutor, sb sqlbuilder.Selector) Selector {
	return &selector{
		sb:   sb,
		exec: exec,
	}
}
