package sqlbuilder

import (
	"reflect"
	"testing"

	"github.com/luanruisong/borm/reflectx"
)

func TestTableName(t *testing.T) {
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

	t.Log(TableName(s))
	t.Log(TableName(s1))
	t.Log(TableName(s2))
}

func TestColumnName(t *testing.T) {
	s := struct {
		AaaAa string `db:"a"`
		BbbBb int
		CccCc float64
	}{}

	reflectx.StructRange(s, func(f reflect.StructField, v reflect.Value) error {
		t.Log(ColumnName(f))
		return nil
	})

}

func TestStructToWhere(t *testing.T) {
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

	t.Log(StructToWhere(s))
	t.Log(StructToWhere(s1))
	t.Log(StructToWhere(s2))
}

func TestIsPk(t *testing.T) {
	s := struct {
		AaaAa string `db:"a"`
		BbbBb int
		CccCc float64 `pk:"1"`
	}{}

	reflectx.StructRange(s, func(f reflect.StructField, v reflect.Value) error {
		t.Log(f.Name, IsPk(f))
		return nil
	})

}
