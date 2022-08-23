package spec

// DeleteStmt represents a delete statement.
type DeleteStmt struct {
	// Action represents the db action.
	Action Action
	// Comment represents a sql comment.
	Comment
	// From represents the operation table name, do not support multiple tables.
	From string
	// Limit represents the limit clause.
	Limit *Limit
	// OrderBy represents the order by clause.
	OrderBy ByItems
	// SQL represents the original sql text.
	SQL string
	// Where represents the where clause.
	Where *Clause

	// the below data are from table
	// FromInfo is the table info which is convert from From.
	FromInfo *Table
}

func (d *DeleteStmt) SQLText() string {
	return d.SQL
}

func (d *DeleteStmt) TableName() string {
	return d.From
}

func (d *DeleteStmt) validate() (map[string]string, error) {
	return map[string]string{
		d.FuncName: d.OriginText,
	}, d.Comment.validate()
}

func (d *DeleteStmt) HasArg() bool {
	if d.Limit.IsValid() {
		return true
	}
	if d.OrderBy.IsValid() {
		return true
	}
	if d.Where.IsValid() {
		return true
	}
	return false
}
