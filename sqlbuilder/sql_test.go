package sqlbuilder

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	sql := InsertInto("table")
	sql.Set("a", 1).
		Set("b", true).
		Set("c", "true")
	t.Log(sql.Sql())
	t.Log(sql.Args())
}

func TestUpdate(t *testing.T) {
	sql := UpdateFrom("table")
	sql.Set("a", 1).
		Set("b", true).
		Set("c", "true")

	sql.Where("id=?", 1)
	t.Log(sql.Sql())
	t.Log(sql.Args())
}

func TestDelete(t *testing.T) {
	sql := DeleteFrom("table")

	sql.Where("id=?", 1)
	t.Log(sql.Sql())
	t.Log(sql.Args())
}

func TestSelect(t *testing.T) {
	sql := SelectFrom("table")
	sql.Select("a", "b", "c").
		Where("id=? and time = ?", 1, time.Now()).
		GroupBy("g1").
		OrderBy("o1").
		Limit(10).
		Offset(200)
	t.Log(sql.Sql())
	t.Log(sql.Args())
}
