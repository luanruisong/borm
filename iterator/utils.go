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
	var err error
	objT := reflect.New(itemT)
	values := make([]interface{}, len(columns))
	fieldMap, _ := reflectx.StructMap(objT.Interface())
	tmpMap := make(map[string]reflect.Value)
	for i, k := range columns {
		f, ok := fieldMap[k]
		if !ok {
			values[i] = new(interface{})
			continue
		}
		values[i] = f.Addr().Interface()
		if ok {
			switch values[i].(type) {
			case time.Time, *time.Time:
				tmpValue := reflect.New(reflect.TypeOf(""))
				values[i] = tmpValue.Interface()
				tmpMap[k] = tmpValue
			}
		}
	}
	err = rows.Scan(values...)
	if err == nil {
		for k, v := range tmpMap {
			curr := v.Elem().String()
			t, _ := time.Parse("2006-01-02 15:04:05", curr)
			currV := reflect.ValueOf(t)
			fieldMap[k].Set(currV)
		}
	}
	return objT, err
}
