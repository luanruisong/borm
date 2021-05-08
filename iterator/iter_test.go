package iterator

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/luanruisong/borm/sqlbuilder"
)

type TestStruct struct {
	Id      uint64
	TString string
	TBool   bool
	TTime   time.Time
}

func open() *sql.DB {
	dsn := "stt_weibo:stt_weibo@tcp(172.16.1.112:3306)/test?loc=Asia%2FShanghai&charset=utf8mb4"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func fmtStr(i interface{}) string {
	b, _ := json.MarshalIndent(i, "", "  ")
	return string(b)
}

func TestIterator_All(t *testing.T) {

	db := open()
	defer db.Close()
	selector := sqlbuilder.SelectFrom("borm_test").Limit(10)

	rows, err := db.Query(selector.Sql(), selector.Args()...)
	if err != nil {
		t.Error(err.Error())
		return
	}

	ite := New(rows)
	res := make([]TestStruct, 0)
	err = ite.All(&res)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(fmtStr(res))
}

func TestIterator_One(t *testing.T) {
	db := open()
	defer db.Close()
	selector := sqlbuilder.SelectFrom("test_struct").Limit(1)
	rows, err := db.Query(selector.Sql(), selector.Args()...)
	if err != nil {
		t.Error(err.Error())
		return
	}
	ite := New(rows)
	res := TestStruct{}
	if err = ite.One(&res); err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(fmtStr(res))
}
