package spec

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/anqiansong/sqlgen/internal/buffer"
	"github.com/anqiansong/sqlgen/internal/set"
)

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

	// the below data are from table
	// ColumnInfo are the column info which are convert from Column.
	ColumnInfo Column
	// TableInfo is the table info.
	TableInfo *Table

	// the below data are from stmt
	// Comment represents a sql comment.
	Comment Comment
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

// Range calls f for each clause, the name of Parameter will add a repeated number suffix to
// avoid duplication Parameter.
func (p Parameters) Range(fn func(p Parameter)) {
	var m = map[Parameter]int{}
	var deduplication Parameters
	for _, v := range p {
		if i, ok := m[v]; ok {
			v.Column = fmt.Sprintf("%s%d", v.Column, i+1)
			m[v] = i + 1
			deduplication = append(deduplication, v)
			continue
		}
		m[v] = 0
		deduplication = append(deduplication, v)
	}

	for _, v := range deduplication {
		fn(v)
	}
}

// IsValid returns true if the statement is valid.
func (c *Clause) IsValid() bool {
	if c == nil {
		return false
	}

	return c.Column != "" || c.OP != 0 || c.Left != nil || c.Right != nil
}

// SQL returns the clause condition strings.
func (c *Clause) SQL() (string, error) {
	sql, _, err := c.marshal()
	return sql, err
}

// ParameterStructure returns the parameter type structure.
func (c *Clause) ParameterStructure(identifier string) (string, error) {
	_, parameters, err := c.marshal()
	if err != nil {
		return "", err
	}

	var writer = buffer.New()
	writer.Write(`type %s struct {`, c.ParameterStructureName(identifier))
	parameters.Range(func(v Parameter) {
		writer.Write("%s %s", v.Column, v.Type)
	})
	writer.Write(`}`)

	return writer.String(), nil
}

// ParameterStructureName returns the parameter structure name.
func (c *Clause) ParameterStructureName(identifier string) string {
	return strcase.ToCamel(fmt.Sprintf("%s%sParameter", c.Comment.FuncName, identifier))
}

// ParameterThirdImports returns the third package imports.
func (c *Clause) ParameterThirdImports() (string, error) {
	_, parameters, err := c.marshal()
	if err != nil {
		return "", err
	}
	var thirdPkgSet = set.From()
	parameters.Range(func(v Parameter) {
		if len(v.ThirdPkg) == 0 {
			return
		}
		thirdPkgSet.Add(v.ThirdPkg)
	})

	return strings.Join(thirdPkgSet.String(), "\n"), nil
}

// Parameters returns the parameter variables.
func (c *Clause) Parameters(pkg string) (string, error) {
	_, parameters, err := c.marshal()
	if err != nil {
		return "", err
	}
	var list []string
	parameters.Range(func(v Parameter) {
		list = append(list, fmt.Sprintf("%s.%s", pkg, v.Column))
	})

	return strings.Join(list, ", "), nil
}

func (c *Clause) marshal() (sql string, parameters Parameters, err error) {
	parameters = []Parameter{}
	if c == nil {
		return
	}

	switch c.OP {
	case And, Or:
		leftSQL, leftParameter, err := c.Left.marshal()
		if err != nil {
			return "", nil, err
		}

		rightSQL, rightParameter, err := c.Right.marshal()
		if err != nil {
			return "", nil, err
		}

		parameters = append(parameters, leftParameter...)
		parameters = append(parameters, rightParameter...)
		var sqlList []string
		if len(leftSQL) > 0 {
			sqlList = append(sqlList, leftSQL)
		}
		if len(leftSQL) > 0 {
			sqlList = append(sqlList, rightSQL)
		}

		sql = strings.Join(sqlList, Operator[c.OP])
	case EQ, GE, GT, In, LE, LT, Like, NE, Not, NotIn, NotLike:
		sql = fmt.Sprintf("%s %s ?", c.Column, Operator[c.OP])
		goType, thirdPkg, err := c.ColumnInfo.Go()
		if err != nil {
			return "", nil, err
		}

		parameters = append(parameters, Parameter{
			Column:   c.Column,
			Type:     goType,
			ThirdPkg: thirdPkg,
		})
	case Between, NotBetween:
		sql = fmt.Sprintf("%s %s ? AND ?", c.Column, Operator[c.OP])
		goType, thirdPkg, err := c.ColumnInfo.Go()
		if err != nil {
			return "", nil, err
		}
		parameters = append(parameters, Parameter{
			Column:   fmt.Sprintf("%s%sStart", c.Column, Operator[c.OP]),
			Type:     goType,
			ThirdPkg: thirdPkg,
		}, Parameter{
			Column:   fmt.Sprintf("%s%sEnd", c.Column, Operator[c.OP]),
			Type:     goType,
			ThirdPkg: thirdPkg,
		})
	default:
		// ignores 'case'
	}
	return
}
