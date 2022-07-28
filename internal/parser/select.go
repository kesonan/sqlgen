package parser

import (
	"fmt"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/opcode"
	"github.com/pingcap/parser/test_driver"

	"github.com/anqiansong/sqlgen/internal/set"
	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseSelect(stmt *ast.SelectStmt) (*spec.SelectStmt, error) {
	var text = stmt.Text()
	var ret spec.SelectStmt

	tableName, err := parseTableRefsClause(stmt.From)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	if stmt.Where != nil {
		where, err := parseExprNode(stmt.Where)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		ret.Where = where
	}

	if stmt.GroupBy != nil {
		groupBy, err := parseGroupBy(stmt.GroupBy)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		ret.GroupBy = groupBy
	}

	if stmt.Having != nil {
		having, err := parseHaving(stmt.Having)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		ret.Having = having
	}

	if stmt.OrderBy != nil {
		orderBy, err := parseOrderBy(stmt.OrderBy)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		ret.OrderBy = orderBy
	}

	if stmt.Limit != nil {
		limit, err := parseLimit(stmt.Limit)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		ret.Limit = limit
	}

	ret.From = tableName
	ret.Distinct = stmt.Distinct
	ret.Action = spec.ActionRead
	ret.SQL = text
	ret.Columns = parseFieldList(stmt.Fields)

	return &ret, nil
}

func convertOP(in opcode.Op) (spec.OP, error) {
	switch in {
	case opcode.LogicAnd:
		return spec.And, nil
	case opcode.LogicOr:
		return spec.Or, nil
	case opcode.GE:
		return spec.GE, nil
	case opcode.LE:
		return spec.LE, nil
	case opcode.EQ:
		return spec.EQ, nil
	case opcode.NE:
		return spec.NE, nil
	case opcode.LT:
		return spec.LT, nil
	case opcode.GT:
		return spec.GT, nil
	case opcode.Not:
		return spec.Not, nil
	case opcode.In:
		return spec.In, nil
	case opcode.Like:
		return spec.Like, nil
	case opcode.Case:
		return spec.Case, nil
	default:
		return 0, fmt.Errorf("unsupported opcode %s", in)
	}
}
func parseExprNode(node ast.ExprNode) (*spec.Clause, error) {
	if node == nil {
		return nil, errorInvalidExprNode
	}

	var clause spec.Clause
	switch v := node.(type) {
	case *ast.BinaryOperationExpr:
		op, err := convertOP(v.Op)
		if err != nil {
			return nil, err
		}

		leftClause, err := parseExprNode(v.L)
		if err != nil {
			return nil, err
		}

		rightClause, err := parseExprNode(v.R)
		if err != nil {
			return nil, err
		}

		clause.OP = op
		if leftClause.IsValid() {
			if leftClause.OP == spec.ColumnOP {
				clause.Column = leftClause.Column
			} else {
				clause.Left = leftClause
			}
		}

		if rightClause.IsValid() {
			if rightClause.OP == spec.ColumnOP {
				clause.Column = rightClause.Column
			} else {
				clause.Right = rightClause
			}
		}
	case *ast.ColumnNameExpr:
		colName, err := parseColumn(v.Name)
		if err != nil {
			return nil, err
		}

		if len(colName) > 0 {
			clause.OP = spec.ColumnOP
			clause.Column = colName
		}
	case *test_driver.ValueExpr, *test_driver.ParamMarkerExpr:
		// ignores it
	case *ast.ParenthesesExpr:
		c, err := parseExprNode(v.Expr)
		if err != nil {
			return nil, err
		}

		clause.OP = spec.Parentheses
		if c.IsValid() && c.OP == spec.ColumnOP {
			clause.Column = c.Column
		} else {
			clause.Left = c
		}
	case *ast.PatternInExpr:
		var inOP = spec.In
		if v.Not {
			inOP = spec.NotIn
		}

		c, err := parseExprNode(v.Expr)
		if err != nil {
			return nil, err
		}

		clause.OP = inOP
		if c.IsValid() && c.OP == spec.ColumnOP {
			clause.Column = c.Column
		} else {
			clause.Left = c
		}
	case *ast.PatternLikeExpr:
		var likeOP = spec.Like
		if v.Not {
			likeOP = spec.NotLike
		}

		c, err := parseExprNode(v.Expr)
		if err != nil {
			return nil, err
		}

		clause.OP = likeOP
		if c.IsValid() && c.OP == spec.ColumnOP {
			clause.Column = c.Column
		} else {
			clause.Left = c
		}
	case *ast.BetweenExpr:
		var betweenOP = spec.Between
		if v.Not {
			betweenOP = spec.NotBetween
		}

		c, err := parseExprNode(v.Expr)
		if err != nil {
			return nil, err
		}

		clause.OP = betweenOP
		if c.IsValid() && c.OP == spec.ColumnOP {
			clause.Column = c.Column
		} else {
			clause.Left = c
		}
	default:
		return nil, errorInvalidExpr
	}

	return &clause, nil
}

func parseGroupBy(groupBy *ast.GroupByClause) ([]string, error) {
	var columnSet = set.From()
	var groupByItem = groupBy.Items
	for _, item := range groupByItem {
		clause, err := parseExprNode(item.Expr)
		if err != nil {
			return nil, err
		}

		columnSet.AddStringList(getAllColumns(clause))
	}

	return columnSet.String(), nil
}

func getAllColumns(clause *spec.Clause) []string {
	var columnSet = set.From()
	if len(clause.Column) > 0 {
		columnSet.Add(clause.Column)
	}
	if clause.Left != nil {
		columnSet.AddStringList(getAllColumns(clause.Left))
	}
	if clause.Right != nil {
		columnSet.AddStringList(getAllColumns(clause.Right))
	}

	return columnSet.String()
}

func parseHaving(having *ast.HavingClause) (*spec.Clause, error) {
	if having == nil {
		return nil, errorMissingHaving
	}

	return parseExprNode(having.Expr)
}

func parseOrderBy(orderBy *ast.OrderByClause) ([]*spec.ByItem, error) {
	var byItem = orderBy.Items
	var ret []*spec.ByItem
	for _, item := range byItem {
		clauses, err := parseExprNode(item.Expr)
		if err != nil {
			return nil, err
		}

		var allColumns = getAllColumns(clauses)
		for _, column := range allColumns {
			ret = append(ret, &spec.ByItem{
				Column: column,
				Desc:   item.Desc,
			})
		}

	}

	return ret, nil
}

func parseLimit(limit *ast.Limit) (*spec.Limit, error) {
	var count = limit.Count
	var offset = limit.Offset
	var ret spec.Limit
	parseValue := func(node ast.ExprNode) (int64, error) {
		value, ok := node.(*test_driver.ValueExpr)
		if !ok {
			return 0, errorUnsupportedLimitExpr
		}
		return value.Datum.GetInt64(), nil
	}

	if count != nil {
		count, err := parseValue(count)
		if err != nil {
			return nil, err
		}
		ret.Count = count
	}

	if offset != nil {
		offset, err := parseValue(offset)
		if err != nil {
			return nil, err
		}
		ret.Offset = offset
	}

	return &ret, nil
}

func parseFieldList(fieldList *ast.FieldList) []string {
	if fieldList == nil {
		return nil
	}

	var columnSet = set.From()
	for _, f := range fieldList.Fields {
		if f.WildCard != nil {
			return []string{spec.WildCard}
		}
		columnSet.Add(f.Text())
	}

	return columnSet.String()
}

func parseColumns(cols []*ast.ColumnName) ([]string, error) {
	var columnSet = set.From()
	for _, col := range cols {
		colName, err := parseColumn(col)
		if err != nil {
			return nil, err
		}

		if colName != "" {
			columnSet.Add(colName)
		}
	}

	return columnSet.String(), nil
}

func parseColumn(col *ast.ColumnName) (string, error) {
	if col == nil {
		return "", nil
	}
	if col.Table.String() != "" {
		return "", errorTableRefer
	}

	return col.Name.O, nil
}

func parseResultSetNode(node ast.ResultSetNode) (string, error) {
	switch t := node.(type) {
	case *ast.TableSource:
		var source = t.Source
		return parseResultSetNode(source)
	case *ast.TableName:
		return t.Name.String(), nil
	case *ast.SelectStmt:
		return "", errorUnsupportedNestedQuery
	case *ast.SetOprStmt:
		return "", errorUnsupportedUnionQuery
	case *ast.SubqueryExpr:
		return "", errorUnsupportedSubQuery
	default:
		return "", errorUnsupportedTableStyle
	}
}
