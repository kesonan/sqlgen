package spec

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
