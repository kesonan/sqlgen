package spec

const (
	Invalid OP = iota
	ColumnOP
	And
	Between
	Case
	EQ
	Or
	GE
	GT
	In
	LE
	LT
	Like
	NE
	Not
	NotBetween
	NotIn
	NotLike
	Parentheses
)

// OP is opcode type.
type OP int
