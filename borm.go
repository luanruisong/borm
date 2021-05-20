package borm

import (
	"errors"
	"reflect"

	"github.com/luanruisong/borm/db"
)

func New(conn db.Connector) (db.DB, error) {
	if conn == nil || reflect.ValueOf(conn).IsNil() {
		return nil, errors.New("nil connector")
	}
	return db.NewDB(conn.DriverName(), conn.ConnStr())
}
