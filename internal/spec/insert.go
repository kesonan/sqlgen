package spec

// InsertStmt represents a insert statement.
type InsertStmt struct {
	// Action represents the db action.
	Action Action
	// Columns represents the operation columns.
	Columns []string
	// Comment represents a sql comment.
	Comment
	// SQL represents the original sql text.
	SQL string
	// Table represents the operation table name, do not support multiple tables.
	Table string

	// the below data are from table
	// ColumnInfo are the column info which are convert from Columns.
	ColumnInfo Columns
	// TableInfo is the table info which is convert from Table.
	TableInfo *Table
}

func (i *InsertStmt) SQLText() string {
	return i.SQL
}

func (i *InsertStmt) TableName() string {
	return i.Table
}

func (i *InsertStmt) validate() (map[string]string, error) {
	return map[string]string{
		i.FuncName: i.OriginText,
	}, i.Comment.validate()
}

func (i *InsertStmt) HasArg() bool {
	return false
}
