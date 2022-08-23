package spec

// UpdateStmt represents a update statement.
type UpdateStmt struct {
	// Action represents the db action.
	Action Action
	// Columns represents the operation columns.
	Columns []string
	// Comment represents a sql comment.
	Comment
	// Limit represents the limit clause.
	Limit *Limit
	// OrderBy represents the order by clause.
	OrderBy ByItems
	// SQL represents the original sql text.
	SQL string
	// Table represents the operation table name, do not support multiple tables.
	Table string
	// Where represents the where clause.
	Where *Clause

	// the below data are from table
	// ColumnInfo are the column info which are convert from Columns.
	ColumnInfo Columns
	// TableInfo is the table info which is convert from Table.
	TableInfo *Table
}

func (u *UpdateStmt) SQLText() string {
	return u.SQL
}

func (u *UpdateStmt) TableName() string {
	return u.Table
}

func (u *UpdateStmt) validate() (map[string]string, error) {
	return map[string]string{
		u.FuncName: u.OriginText,
	}, u.Comment.validate()
}

func (u *UpdateStmt) HasArg() bool {
	if u.Limit.IsValid() {
		return true
	}
	if u.OrderBy.IsValid() {
		return true
	}
	if u.Where.IsValid() {
		return true
	}
	return false
}
