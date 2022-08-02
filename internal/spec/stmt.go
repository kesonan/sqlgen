package spec

// WildCard is a wildcard column.
const WildCard = "*"

var _ DML = (*InsertStmt)(nil)
var _ DML = (*UpdateStmt)(nil)
var _ DML = (*SelectStmt)(nil)
var _ DML = (*DeleteStmt)(nil)

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
	ColumnInfo []Column
	// TableInfo is the table info which is convert from Table.
	TableInfo *Table
}

// SelectStmt represents a select statement.
type SelectStmt struct {
	// Action represents the db action.
	Action Action
	// Columns represents the operation columns.
	Columns []string
	// Comment represents a sql comment.
	Comment
	// Distinct represents the select distinct flag.
	Distinct bool
	// From represents the operation table name, do not support multiple tables.
	From string
	// GroupBy represents the group by clause.
	GroupBy ByItems
	// Having represents the having clause.
	Having *Clause
	// Limit represents the limit clause.
	Limit *Limit
	// OrderBy represents the order by clause.
	OrderBy ByItems
	// SQL represents the original sql text.
	SQL string
	// Where represents the where clause.
	Where *Clause

	// the below data are from table
	// ColumnInfo are the column info which are convert from Columns.
	ColumnInfo []Column
	// FromInfo is the table info which is convert from From.
	FromInfo *Table
}

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
	ColumnInfo []Column
	// TableInfo is the table info which is convert from Table.
	TableInfo *Table
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

func (i *InsertStmt) SQLText() string {
	return i.SQL
}

func (i *InsertStmt) TableName() string {
	return i.Table
}

func (i *InsertStmt) validate() error {
	return i.Comment.validate()
}

func (s *SelectStmt) SQLText() string {
	return s.SQL
}

func (s *SelectStmt) TableName() string {
	return s.From
}

func (s *SelectStmt) validate() error {
	return s.Comment.validate()
}

func (d *DeleteStmt) SQLText() string {
	return d.SQL
}

func (d *DeleteStmt) TableName() string {
	return d.From
}

func (d *DeleteStmt) validate() error {
	return d.Comment.validate()
}

func (u *UpdateStmt) SQLText() string {
	return u.SQL
}

func (u *UpdateStmt) TableName() string {
	return u.Table
}

func (u *UpdateStmt) validate() error {
	return u.Comment.validate()
}
