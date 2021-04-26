package iterator

import (
	"database/sql"
	"reflect"
	"sync"
)

type (
	iterator struct {
		cursor *sql.Rows
		once   sync.Once
		err    error
	}
	Iterator interface {
		All(i interface{}) error
		One(i interface{}) error
	}
)

func (ite *iterator) close() {
	ite.once.Do(func() {
		if ite.cursor != nil {
			ite.cursor.Close()
		}
	})
}
func (ite *iterator) All(dst interface{}) error {
	dstv := reflect.ValueOf(dst)
	if dstv.IsNil() || dstv.Kind() != reflect.Ptr {
		return ErrNotPtr
	}
	if dstv.Elem().Kind() != reflect.Slice {
		return ErrNotSlice
	}
	defer ite.close()
	rows := ite.cursor
	slicev := dstv.Elem()
	itemT := slicev.Type().Elem()
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	reset(dst)
	for rows.Next() {
		item, err := fetchResult(rows, itemT, columns)
		if err != nil {
			return err
		}
		if itemT.Kind() == reflect.Ptr {
			slicev = reflect.Append(slicev, item)
		} else {
			slicev = reflect.Append(slicev, reflect.Indirect(item))
		}
	}
	dstv.Elem().Set(slicev)
	return rows.Err()
}

func (ite *iterator) One(dst interface{}) error {
	dstv := reflect.ValueOf(dst)
	if dstv.IsNil() || dstv.Kind() != reflect.Ptr {
		return ErrNotPtr
	}
	itemV := dstv.Elem()
	reset(dst)
	defer ite.close()
	rows := ite.cursor
	itemT := dstv.Type().Elem()
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	if !rows.Next() {
		return ErrNoMoreRows
	}
	item, err := fetchResult(rows, itemT, columns)
	if err != nil {
		return err
	}
	if itemT.Kind() == reflect.Ptr {
		itemV.Set(item)
	} else {
		itemV.Set(reflect.Indirect(item))
	}
	return rows.Err()

}

func New(rs *sql.Rows) *iterator {
	return &iterator{rs, sync.Once{}, nil}
}
