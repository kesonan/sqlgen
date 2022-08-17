package spec

import (
	_ "embed"
)

// WildCard is a wildcard column.
const WildCard = "*"

var _ DML = (*InsertStmt)(nil)
var _ DML = (*UpdateStmt)(nil)
var _ DML = (*SelectStmt)(nil)
var _ DML = (*DeleteStmt)(nil)

type Fields []Field

// Field represents a select filed.
type Field struct {
	ASName        string
	ColumnName    string
	TP            byte
	AggregateCall bool
}

// Limit represents a limit clause.
type Limit struct {
	// Count represents the limit count.
	Count int64
	// Offset represents the limit offset.
	Offset int64

	// the below data are from table
	// TableInfo is the table info.
	TableInfo *Table

	// the below data are from stmt
	// Comment represents a sql comment.
	Comment Comment
}

func (f Field) Name() string {
	if len(f.ASName) > 0 {
		return f.ASName
	}
	return f.ColumnName
}
