package sql

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"

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

func (c customScanner) getRowsElem(rows *sql.Rows, v interface{}) error {
	tp := reflect.TypeOf(v)
	if tp.Kind() != reflect.Pointer {
		return errors.New("expected a pointer")
	}
	sliceTp := tp.Elem()
	if sliceTp.Kind() != reflect.Slice {
		return errors.New("expected a slice")
	}

	sliceValue := reflect.Indirect(reflect.ValueOf(v))
	itemType := sliceTp.Elem()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	for rows.Next() {
		item := reflect.New(itemType.Elem()).Elem()
		dest := structPointers(item.Elem(), cols)

		err := rows.Scan(dest...)
		if err != nil {
			return err
		}
		sliceValue.Set(reflect.Append(sliceValue, item))
	}

	return rows.Err()
}

func fieldByName(v reflect.Value, name string) reflect.Value {
	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag, ok := typ.Field(i).Tag.Lookup("db")
		if ok && tag == name {
			return v.Field(i)
		}
	}

	return v.FieldByName(strings.Title(name))
}

func structPointers(stct reflect.Value, cols []string) []interface{} {
	pointers := make([]interface{}, 0, len(cols))
	for _, colName := range cols {
		fieldVal := fieldByName(stct, colName)
		if !fieldVal.IsValid() || !fieldVal.CanSet() {
			var nothing interface{}
			pointers = append(pointers, &nothing)
			continue
		}

		pointers = append(pointers, fieldVal.Addr().Interface())
	}
	return pointers
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
