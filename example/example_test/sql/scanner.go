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
	var elem reflect.Value
	value, ok := v.(reflect.Value)
	if !ok {
		elem = value.Elem()
		value = reflect.ValueOf(v)
	} else {
		elem = value
	}

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

// getRowsElem is inspired by https://github.com/zeromicro/go-zero/blob/8ed22eafdda04c4526164450d7c13c2f4b0f076c/core/stores/sqlx/orm.go#L163
func (c customScanner) getRowsElem(rows *sql.Rows, v interface{}) error {
	valueOf := reflect.ValueOf(v)
	if valueOf.Kind() != reflect.Ptr {
		return errors.New("expect a pointer")
	}

	typeOf := reflect.TypeOf(v)
	sliceTypeOf := typeOf.Elem()
	sliceValueOf := valueOf.Elem()

	if sliceTypeOf.Kind() != reflect.Slice {
		return errors.New("expect a slice")
	}
	if !sliceValueOf.CanSet() {
		return errors.New("expect a settable slice")
	}
	isASlicePointer := sliceTypeOf.Elem().Kind() == reflect.Ptr

	var itemReceiver reflect.Type
	itemType := sliceTypeOf.Elem()
	if itemType.Kind() == reflect.Ptr {
		itemReceiver = itemType.Elem()
	} else {
		itemReceiver = itemType
	}
	if itemReceiver.Kind() != reflect.Struct {
		return errors.New("expect a struct")
	}

	for rows.Next() {
		value := reflect.New(itemReceiver)
		dest, err := c.getRowElem(value)
		if err != nil {
			return err
		}

		err = rows.Scan(dest...)
		if err != nil {
			return err
		}

		if isASlicePointer {
			sliceValueOf.Set(reflect.Append(sliceValueOf, value))
		} else {
			sliceValueOf.Set(reflect.Append(sliceValueOf, reflect.Indirect(sliceValueOf)))
		}
	}

	return nil
}

func (c customScanner) ScanRow(row *sql.Row, v interface{}) error {
	dest, err := c.getRowElem(v)
	if err != nil {
		return err
	}
	return row.Scan(dest...)
}

func (c customScanner) ScanRows(rows *sql.Rows, v interface{}) error {
	//return scan.Rows(v, rows)
	return c.getRowsElem(rows, v)
}

func getScanner() model.Scanner {
	return customScanner{}
}
