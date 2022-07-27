package spec

// WildCard is a wildcard column.
const WildCard = "*"

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

type Limit struct {
	Count  bool
	Offset bool
}

func (c *Clause) IsValid() bool {
	return c.Column != "" || c.OP != 0 || c.Left != nil || c.Right != nil
}

func (i *InsertStmt) SQLText() string {
	return i.SQL
}

func (s *SelectStmt) SQLText() string {
	return s.SQL
}

func (d *DeleteStmt) SQLText() string {
	return d.SQL
}

func (u *UpdateStmt) SQLText() string {
	return u.SQL
}
