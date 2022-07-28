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
	// SQL represents the original sql text.
	SQL string
	// Table represents the operation table name, do not support multiple tables.
	Table string
}

// SelectStmt represents a select statement.
type SelectStmt struct {
	// Action represents the db action.
	Action Action
	// Columns represents the operation columns.
	Columns []string
	// Distinct represents the select distinct flag.
	Distinct bool
	// From represents the operation table name, do not support multiple tables.
	From string
	// GroupBy represents the group by clause.
	GroupBy []string
	// Having represents the having clause.
	Having *Clause
	// Limit represents the limit clause.
	Limit *Limit
	// OrderBy represents the order by clause.
	OrderBy []*ByItem
	// SQL represents the original sql text.
	SQL string
	// Where represents the where clause.
	Where *Clause
}

// DeleteStmt represents a delete statement.
type DeleteStmt struct {
	// Action represents the db action.
	Action Action
	// From represents the operation table name, do not support multiple tables.
	From string
	// Limit represents the limit clause.
	Limit *Limit
	// OrderBy represents the order by clause.
	OrderBy []*ByItem
	// SQL represents the original sql text.
	SQL string
	// Where represents the where clause.
	Where *Clause
}

// UpdateStmt represents a update statement.
type UpdateStmt struct {
	// Action represents the db action.
	Action Action
	// Columns represents the operation columns.
	Columns []string
	// Limit represents the limit clause.
	Limit *Limit
	// OrderBy represents the order by clause.
	OrderBy []*ByItem
	// SQL represents the original sql text.
	SQL string
	// Table represents the operation table name, do not support multiple tables.
	Table string
	// Where represents the where clause.
	Where *Clause
}

// ByItem represents an order-by or group-by item.
type ByItem struct {
	Column string
	Desc   bool
}

// Clause represents a where clause, having clause.
type Clause struct {
	// Column represents the column name.
	Column string
	// Left represents the left expr.
	Left *Clause
	// Right represents the right expr.
	Right *Clause
	// OP represents the operator.
	OP OP
}

// Limit represents a limit clause.
type Limit struct {
	// Count represents the limit count.
	Count int64
	// Offset represents the limit offset.
	Offset int64
}

// IsValid returns true if the statement is valid.
func (c *Clause) IsValid() bool {
	return c.Column != "" || c.OP != 0 || c.Left != nil || c.Right != nil
}

func (i *InsertStmt) SQLText() string {
	return i.SQL
}

func (i *InsertStmt) validate() error {
	return nil
}

func (s *SelectStmt) SQLText() string {
	return s.SQL
}

func (s *SelectStmt) validate() error {
	return nil
}

func (d *DeleteStmt) SQLText() string {
	return d.SQL
}

func (d *DeleteStmt) validate() error {
	return nil
}

func (u *UpdateStmt) SQLText() string {
	return u.SQL
}

func (u *UpdateStmt) validate() error {
	return nil
}
