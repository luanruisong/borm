package borm

import (
	"errors"

	"github.com/luanruisong/borm/db"
	"github.com/luanruisong/borm/reflectx"
)

func New(conn db.Connector) (db.DB, error) {
	if reflectx.IsNull(conn) {
		return nil, errors.New("nil connector")
	}
	db, err := db.NewDB(conn.DriverName(), conn.ConnStr())
	if err == nil {
		db.SetMaxOpenConns(conn.GetPoolSize())
	}
	return db, err
}
