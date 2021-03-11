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

func TestAutoInsert(t *testing.T) {
	s := struct {
		AaaAa string `db:"a"`
		BbbBb int
		CccCc float64
	}{}

	type TestTable struct {
		AaaAa string `db:"a"`
		BbbBb int
		CccCc float64
		asda  bool
	}

	s1 := TestTable{}
	s2 := TestTable{
		AaaAa: "asdasdasdasdas",
		BbbBb: 200,
		CccCc: 3.14,
		asda:  true,
	}

	autoInsert(t, s)
	autoInsert(t, s1)
	autoInsert(t, s2)

}

func autoInsert(t *testing.T, s interface{}) {
	sql, err := AutoInsert(s)
	if err != nil {
		t.Error(err.Error())
		return
	} else {
		t.Log(sql.Sql(), sql.Args())
	}
}

func TestAutoUpdate(t *testing.T) {
	type TestTable struct {
		AaaAa string `db:"a"`
		BbbBb int
		CccCc float64
		asda  bool
		Id    string `dbw:"gt"`
		Ts    int64  `dbw:"lt"`
	}
	s := TestTable{
		AaaAa: "asdasdasdasdas",
		BbbBb: 200,
		CccCc: 3.14,
		asda:  true,
		Id:    "10000",
		Ts:    time.Now().Unix(),
	}
	sql, err := AutoUpdate(s)
	if err != nil {
		t.Error(err.Error())
		return
	} else {
		t.Log(sql.Sql(), sql.Args())
	}
}

type TestSelectTable struct {
	AaaAa string `db:"a"`
	BbbBb int
	CccCc float64
	asda  bool
	Id    string `dbw:"gt"`
	Ts    int64  `dbw:"lt"`
}

func (TestSelectTable) TableName() string {
	return "asd"
}
func TestAutoSelect(t *testing.T) {

	s := TestSelectTable{
		AaaAa: "asdasdasdasdas",
		BbbBb: 200,
		CccCc: 3.14,
		asda:  true,
		Id:    "10000",
		Ts:    time.Now().Unix(),
	}
	sql, err := AutoSelect(s)
	if err != nil {
		t.Error(err.Error())
		return
	} else {
		t.Log(sql.Sql(), sql.Args())
	}
}
