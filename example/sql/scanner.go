package model

import "database/sql"

type Scanner interface {
	ScanRow(rows *sql.Rows, v interface{}) error
	ScanRows(rows *sql.Rows, v interface{}) error
	ColumnMapper(colName string) string
	TagKey() string
}
