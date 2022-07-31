package spec

import (
	"fmt"
	"strings"
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
}

type Parameter struct {
	Column string
	Type   string
}

// IsValid returns true if the statement is valid.
func (c *Clause) IsValid() bool {
	if c == nil {
		return false
	}

	return c.Column != "" || c.OP != 0 || c.Left != nil || c.Right != nil
}

func (c *Clause) SQL() (string, error) {
	panic("implement me")
}

func (c *Clause) ParameterStructure() (string, error) {
	panic("implement me")
}

func (c *Clause) Parameters(pkg string) (string, error) {
	panic("implement me")
}

func (c *Clause) marshal() (sql string, parameters []Parameter, err error) {
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
		goType, err := c.ColumnInfo.Go()
		if err != nil {
			return "", nil, err
		}

		parameters = append(parameters, Parameter{
			Column: c.Column,
			Type:   goType,
		})
	case Between, NotBetween:
		sql = fmt.Sprintf("%s %s ? AND ?", c.Column, Operator[c.OP])
		goType, err := c.ColumnInfo.Go()
		if err != nil {
			return "", nil, err
		}
		parameters = append(parameters, Parameter{
			Column: fmt.Sprintf("%s%sStart", c.Column, Operator[c.OP]),
			Type:   goType,
		}, Parameter{
			Column: fmt.Sprintf("%s%sEnd", c.Column, Operator[c.OP]),
			Type:   goType,
		})
	default:
		// ignores 'case'
	}
	return
}
