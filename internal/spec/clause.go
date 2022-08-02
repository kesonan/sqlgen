package spec

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/anqiansong/sqlgen/internal/buffer"
	"github.com/anqiansong/sqlgen/internal/parameter"
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

// NewParameter returns a new parameter.
func NewParameter(column string, tp string, thirdPkg string) parameter.Parameter {
	return parameter.Parameter{Column: strcase.ToCamel(column), Type: tp, ThirdPkg: thirdPkg}
}

// IsInValid returns true if the statement is valid.
func (c *Clause) IsInValid() bool {
	if c == nil {
		return true
	}

	return !(c.Column != "" || c.OP != 0 || c.Left != nil || c.Right != nil)
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
	for _, v := range parameters {
		writer.Write("%s %s", v.Column, v.Type)
	}

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
	for _, v := range parameters {
		if len(v.ThirdPkg) == 0 {
			continue
		}
		thirdPkgSet.Add(v.ThirdPkg)
	}

	return strings.Join(thirdPkgSet.String(), "\n"), nil
}

// Parameters returns the parameter variables.
func (c *Clause) Parameters(pkg string) (string, error) {
	_, parameters, err := c.marshal()
	if err != nil {
		return "", err
	}
	var list []string
	for _, v := range parameters {
		list = append(list, fmt.Sprintf("%s.%s", pkg, v.Column))
	}

	return strings.Join(list, ", "), nil
}

func (c *Clause) marshal() (sql string, parameters parameter.Parameters, err error) {
	parameters = parameter.Empty
	if c == nil {
		return
	}

	var ps = parameter.New()
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

		ps.Add(leftParameter...)
		ps.Add(rightParameter...)
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
		p, err := c.ColumnInfo.DataType()
		if err != nil {
			return "", nil, err
		}

		ps.Add(p)
	case Between, NotBetween:
		sql = fmt.Sprintf("%s %s ? AND ?", c.Column, Operator[c.OP])
		p, err := c.ColumnInfo.DataType()
		if err != nil {
			return "", nil, err
		}

		ps.Add(
			NewParameter(fmt.Sprintf("%s%sStart", c.Column, OpName[c.OP]), p.Type, p.ThirdPkg),
			NewParameter(fmt.Sprintf("%s%sEnd", c.Column, OpName[c.OP]), p.Type, p.ThirdPkg))
	default:
		// ignores 'case'
	}
	parameters = ps.List()
	return
}
