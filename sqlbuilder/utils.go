package sqlbuilder

import (
	"reflect"

	"github.com/luanruisong/borm/reflectx"
	"github.com/luanruisong/borm/stringx"
)

func TableName(i interface{}) string {
	var tName string
	if t, ok := i.(Table); ok {
		tName = t.TableName()
	} else {
		//未实现接口，取struct名称的蛇形
		tName = stringx.SnakeName(reflectx.StructName(i))
	}
	return tName
}

func IsPk(t reflect.StructField) bool {
	tag := t.Tag.Get("pk")
	if len(tag) == 0 {
		//入无自定义column，取field名称的蛇形
		return false
	}
	return tag == "1"
}
