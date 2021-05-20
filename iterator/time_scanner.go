package iterator

import (
	"database/sql"
	"reflect"
	"time"
)

type (
	TimeScanner struct {
		v      reflect.Value
		format string
	}
)

const defTimeFormat = "2006-01-02 15:04:05"

func (t TimeScanner) Scan(src interface{}) (err error) {
	var (
		curr time.Time
	)
	switch src.(type) {
	case time.Time:
		curr = src.(time.Time)
	case []byte:
		curr, err = time.Parse(t.format, string(src.([]byte)))
		if err != nil {
			return
		}
	case string:
		curr, err = time.Parse(t.format, src.(string))
		if err != nil {
			return
		}
	case int64:
		curr = time.Unix(src.(int64), 0)
	case nil:
		return
	default:
		return ErrTimeScan
	}
	currV := reflect.ValueOf(curr)
	t.v.Set(currV)
	return
}

func NewTimeScanner(v reflect.Value, format string) sql.Scanner {
	return &TimeScanner{
		v:      v,
		format: format,
	}
}
