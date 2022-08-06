package spec

import "fmt"

// Comment represents a sql comment.
type Comment struct {
	// OriginText represents the original sql text.
	OriginText string
	// LineText is the text of the line comment.
	LineText []string
	// FuncNames represents the generated function names.
	FuncName string
}

func (c Comment) validate() error {
	if len(c.FuncName) == 0 {
		return fmt.Errorf("missing func name near %s", c.OriginText)
	}
	return nil
}
