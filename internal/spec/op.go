package spec

const (
	_ OP = iota
	ColumnValue
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

var Operator = []string{
	And:        "AND",
	Between:    "BETWEEN",
	Case:       "CASE",
	EQ:         "=",
	Or:         "OR",
	GE:         ">=",
	GT:         ">",
	In:         "IN",
	LE:         "<=",
	LT:         "<",
	Like:       "LIKE",
	NE:         "!=",
	Not:        "NOT",
	NotBetween: "NOT BETWEEN",
	NotIn:      "NOT IN",
	NotLike:    "NOT LIKE",
}

var OpName = []string{
	And:        "And",
	Between:    "Between",
	Case:       "Case",
	EQ:         "Equal",
	Or:         "Or",
	GE:         "GE",
	GT:         "GT",
	In:         "In",
	LE:         "LE",
	LT:         "LT",
	Like:       "Like",
	NE:         "NE",
	Not:        "Not",
	NotBetween: "NotBetween",
	NotIn:      "NotIn",
	NotLike:    "NotLike",
}
