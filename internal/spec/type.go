package spec

import "github.com/pingcap/parser/mysql"

// Type is the type of the column.
type Type byte

func (t Type) Go() string {
	switch t {
	case mysql.TypeDouble:
		return "string"
	}
	return ""
}
