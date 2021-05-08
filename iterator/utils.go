package iterator

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/luanruisong/borm/reflectx"
)

//reset 重置interface
func reset(data interface{}) error {
	v := reflect.ValueOf(data).Elem()
	t := v.Type()
	var z reflect.Value
	switch v.Kind() {
	case reflect.Slice:
		z = reflect.MakeSlice(t, 0, v.Cap())
	default:
		z = reflect.Zero(t)
	}
	v.Set(z)
	return nil
}

//fetchResult 通过列名抓取响应属性生成一个类型指针
func fetchResult(rows *sql.Rows, itemT reflect.Type, columns []string) (reflect.Value, error) {
	objT := reflect.New(itemT)
	values := make([]interface{}, len(columns))
	fieldMap, _ := reflectx.StructMap(objT.Interface())
	for i, k := range columns {
		f, ok := fieldMap[k]
		if !ok {
			values[i] = new(interface{})
			continue
		}
		curr := f.Addr().Interface()
		switch curr.(type) {
		case time.Time, *time.Time:
			format := defTimeFormat
			if dbFmt := f.Tag.Get("fmt"); len(dbFmt) > 0 {
				format = dbFmt
			}
			values[i] = NewTimeScanner(f.Value, format)
		default:
			values[i] = curr
		}
	}
	return objT, rows.Scan(values...)
}
