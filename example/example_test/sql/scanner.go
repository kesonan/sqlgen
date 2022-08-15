package sql

import (
	"database/sql"
	"errors"
	"reflect"

	model "github.com/anqiansong/sqlgen/example/sql"
)

type customScanner struct {
}

func (c customScanner) getRowElem(v interface{}) ([]interface{}, error) {
	value, ok := v.(reflect.Value)
	if !ok {
		value = reflect.ValueOf(v)
	}
	elem := value.Elem()
	switch elem.Kind() {
	case reflect.Pointer:
		return c.getRowElem(elem.Elem())
	case reflect.Struct:
		var list []interface{}
		for i := 0; i < elem.NumField(); i++ {
			f := elem.Field(i)
			list = append(list, f.Addr().Interface())
		}
		return list, nil
	default:
		return nil, errors.New("expect a struct")
	}
}

func (c customScanner) getRowsElem(v interface{}) ([][]interface{}, error) {
	value := reflect.ValueOf(v)
	elem := value.Elem()
	switch elem.Kind() {
	case reflect.Pointer:
		return c.getRowsElem(elem.Elem())
	case reflect.Slice:
		var list [][]interface{}
		for i := 0; i < elem.NumField(); i++ {
			f := elem.Field(i)
			item := f.Elem()
			rowElem, err := c.getRowsElem(item)
			if err != nil {
				return nil, err
			}

			list = append(list, rowElem...)
		}
		return list, nil
	default:
		return nil, errors.New("expect a struct")
	}
}

func (c customScanner) ScanRow(row *sql.Row, v interface{}) error {
	dest, err := c.getRowElem(v)
	if err != nil {
		return err
	}
	return row.Scan(dest...)
}

func (c customScanner) ScanRows(rows *sql.Rows, v interface{}) error {
	dests, err := c.getRowsElem(v)
	if err != nil {
		return err
	}
	var i int
	for rows.Next() && i < len(dests) {
		dest := dests[i]
		err = rows.Scan(dest)
		if err != nil {
			return err
		}
		i++
	}
	return nil
}

func getScanner() model.Scanner {
	return customScanner{}
}
