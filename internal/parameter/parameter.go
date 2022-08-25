package parameter

import (
	"github.com/anqiansong/sqlgen/internal/set"
	"github.com/anqiansong/sqlgen/internal/stringx"
)

type p struct {
	s *set.ListSet
}

// Parameter represents an original description data for code generation structure info.
type Parameter struct {
	// Column represents a parameter name.
	Column string
	// Type represents a parameter go type.
	Type string
	// ThirdPkg represents a go type which is a third package or go built-in package.
	ThirdPkg string
}

// Parameters returns the parameters.
type Parameters []Parameter

// Empty is a placeholder of Parameters.
var Empty = Parameters{}

func New() *p {
	return &p{s: set.From()}
}

func (p *p) Add(parameter ...Parameter) {
	for _, v := range parameter {
		for {
			if p.s.Exists(v) {
				v.Column = stringx.AutoIncrement(v.Column, 1)
				continue
			}
			p.s.Add(v)
			break
		}
	}
}

func (p *p) List() Parameters {
	var ret Parameters
	p.s.Range(func(v interface{}) {
		ret = append(ret, v.(Parameter))
	})
	return ret
}
