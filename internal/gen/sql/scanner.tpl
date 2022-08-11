package model

import "database/sql"

type Scanner interface {
    ScanRow(row *sql.Row, v interface{}) error
    ScanRows(rows *sql.Rows, v interface{}) error
}