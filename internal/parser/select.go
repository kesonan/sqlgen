package parser

import (
	"fmt"
	"strings"

	"github.com/anqiansong/sqlgen/internal/set"
	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/opcode"
	"github.com/pingcap/parser/test_driver"
)

func parseSelect(stmt *ast.SelectStmt, needFn bool) (*spec.SelectStmt, error) {
	var text = stmt.Text()
	comment, err := parseLineComment(text, needFn)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	sql, err := NewSqlScanner(text).ScanAndTrim()
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	var ret spec.SelectStmt
	ret.Comment = comment
	tableName, err := parseTableRefsClause(stmt.From)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	if stmt.Where != nil {
		where, err := parseExprNode(stmt.Where, tableName, exprTypeWhereClause)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		ret.Where = where
	}

	if stmt.GroupBy != nil {
		groupBy, err := parseGroupBy(stmt.GroupBy, tableName)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		ret.GroupBy = groupBy
	}

	if stmt.Having != nil {
		having, err := parseHaving(stmt.Having, tableName)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		ret.Having = having
	}

	if stmt.OrderBy != nil {
		orderBy, err := parseOrderBy(stmt.OrderBy, tableName)
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
	columnNames, selectFieldSQL, isAllAggregate, err := parseFieldList(stmt.Fields, tableName)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	if isAllAggregate && !ret.Limit.IsValid() {
		ret.Limit = &spec.Limit{
			Count: 1,
		}
	}
	ret.From = tableName
	ret.SelectSQL = fmt.Sprintf("%s", selectFieldSQL)
	ret.Distinct = stmt.Distinct
	ret.Action = spec.ActionRead
	ret.SQL = sql
	ret.Columns = columnNames

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

type exprType int

const (
	_ exprType = iota
	exprTypeColumn
	exprTypeWhereClause
	exprTypeHavingClause
	exprTypeGroupByItems
	exprTypeOrderByItems
)

var exprTypeValue = [...]string{
	exprTypeColumn:       "column expr",
	exprTypeWhereClause:  "where clause",
	exprTypeHavingClause: "having clause",
	exprTypeGroupByItems: "group by items",
	exprTypeOrderByItems: "order by items",
}

func parseExprNode(node ast.ExprNode, table string, exprType exprType) (*spec.Clause, error) {
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

		leftClause, err := parseExprNode(v.L, table, exprType)
		if err != nil {
			return nil, err
		}

		rightClause, err := parseExprNode(v.R, table, exprType)
		if err != nil {
			return nil, err
		}

		clause.OP = op
		if leftClause.IsValid() {
			if leftClause.OP == spec.ColumnValue {
				clause.Column = leftClause.Column
			} else {
				clause.Left = leftClause
			}
		}

		if rightClause.IsValid() {
			if rightClause.OP == spec.ColumnValue {
				clause.Column = rightClause.Column
			} else {
				clause.Right = rightClause
			}
		}
	case *ast.ColumnNameExpr:
		colName, err := parseColumn(v.Name, table)
		if err != nil {
			return nil, err
		}

		if len(colName) > 0 {
			clause.OP = spec.ColumnValue
			clause.Column = colName
		}
	case *test_driver.ValueExpr, *test_driver.ParamMarkerExpr:
		// ignores it
	case *ast.ParenthesesExpr:
		c, err := parseExprNode(v.Expr, table, exprType)
		if err != nil {
			return nil, err
		}

		clause.OP = spec.Parentheses
		if c.IsValid() && c.OP == spec.ColumnValue {
			clause.Column = c.Column
		} else {
			clause.Left = c
		}
	case *ast.PatternInExpr:
		var inOP = spec.In
		if v.Not {
			inOP = spec.NotIn
		}

		c, err := parseExprNode(v.Expr, table, exprType)
		if err != nil {
			return nil, err
		}

		clause.OP = inOP
		if c.IsValid() && c.OP == spec.ColumnValue {
			clause.Column = c.Column
		} else {
			clause.Left = c
		}
	case *ast.PatternLikeExpr:
		var likeOP = spec.Like
		if v.Not {
			likeOP = spec.NotLike
		}

		c, err := parseExprNode(v.Expr, table, exprType)
		if err != nil {
			return nil, err
		}

		clause.OP = likeOP
		if c.IsValid() && c.OP == spec.ColumnValue {
			clause.Column = c.Column
		} else {
			clause.Left = c
		}
	case *ast.BetweenExpr:
		var betweenOP = spec.Between
		if v.Not {
			betweenOP = spec.NotBetween
		}

		c, err := parseExprNode(v.Expr, table, exprType)
		if err != nil {
			return nil, err
		}

		clause.OP = betweenOP
		if c.IsValid() && c.OP == spec.ColumnValue {
			clause.Column = c.Column
		} else {
			clause.Left = c
		}
	default:
		return nil, fmt.Errorf("[%s]: unsupported expr node %T ", exprTypeValue[exprType], v)
	}

	return &clause, nil
}

func parseGroupBy(groupBy *ast.GroupByClause, table string) (spec.ByItems, error) {
	var ret spec.ByItems
	var groupByItem = groupBy.Items
	for _, item := range groupByItem {
		clause, err := parseExprNode(item.Expr, table, exprTypeGroupByItems)
		if err != nil {
			return nil, err
		}

		var allColumns = getAllColumns(clause)
		for _, column := range allColumns {
			ret = append(ret, &spec.ByItem{
				Column: column,
				Desc:   item.Desc,
			})
		}
	}

	return ret, nil
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

func parseHaving(having *ast.HavingClause, table string) (*spec.Clause, error) {
	if having == nil {
		return nil, errorMissingHaving
	}

	return parseExprNode(having.Expr, table, exprTypeHavingClause)
}

func parseOrderBy(orderBy *ast.OrderByClause, table string) (spec.ByItems, error) {
	var byItem = orderBy.Items
	var ret spec.ByItems
	for _, item := range byItem {
		clauses, err := parseExprNode(item.Expr, table, exprTypeOrderByItems)
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

const (
	countMarkerDefaultValue  = 10
	offsetMarkerDefaultValue = 1
)

func parseLimit(limit *ast.Limit) (*spec.Limit, error) {
	var count = limit.Count
	var offset = limit.Offset
	var ret spec.Limit
	parseValue := func(node ast.ExprNode) (int64, error) {
		switch v := node.(type) {
		case *test_driver.ValueExpr:
			return v.Datum.GetInt64(), nil
		case *test_driver.ParamMarkerExpr:
			return 0, errorParamMaker
		default:
			return 0, errorUnsupportedLimitExpr
		}
	}

	if count != nil {
		count, err := parseValue(count)
		if err != nil {
			if err != errorParamMaker {
				return nil, err
			}
			ret.Count = countMarkerDefaultValue
		} else {
			ret.Count = count
		}
	}

	if offset != nil {
		offset, err := parseValue(offset)
		if err != nil {
			if err != errorParamMaker {
				return nil, err
			}
			ret.Offset = offsetMarkerDefaultValue
		} else {
			ret.Offset = offset
		}
	}

	return &ret, nil
}

func parseFieldList(fieldList *ast.FieldList, from string) (spec.Fields, string, bool, error) {
	if fieldList == nil {
		return spec.Fields{}, "", false, nil
	}

	var selectField []string
	var columnSet = set.From()
	var isAllAggregate = true
	for _, f := range fieldList.Fields {
		if f.WildCard != nil {
			wildcardTable := f.WildCard.Table.String()
			if len(wildcardTable) > 0 && wildcardTable != from {
				return nil, "", false, fmt.Errorf("wildcard table %q not match from %q", f.WildCard.Table.String(), from)
			}
			selectField = append(selectField, spec.WildCard)
			columnSet.Add(spec.Field{
				ColumnName: spec.WildCard,
			})
			isAllAggregate = false
			continue
		}

		columnName, funcSql, tp, aggregate, err := parseSelectField(f.Expr, len(f.AsName.String()) > 0, from)
		if err != nil {
			return nil, "", false, err
		}

		if !aggregate {
			isAllAggregate = false
		}

		if len(f.AsName.String()) > 0 {
			funcSql = fmt.Sprintf("%s AS %s", funcSql, f.AsName.String())
		}
		selectField = append(selectField, funcSql)
		columnSet.Add(spec.Field{
			ASName:        f.AsName.String(),
			ColumnName:    columnName,
			TP:            tp,
			AggregateCall: aggregate,
		})
	}

	var fields spec.Fields
	columnSet.Range(func(v interface{}) {
		fields = append(fields, v.(spec.Field))
	})

	return fields, strings.Join(selectField, ", "), isAllAggregate, nil
}

func parseSelectField(node ast.ExprNode, hasAsName bool, table string) (string, string, byte, bool, error) {
	switch v := node.(type) {
	case *ast.ColumnNameExpr:
		columnName, err := parseColumn(v.Name, table)
		if err != nil {
			return "", "", mysql.TypeUnspecified, false, err
		}
		return columnName, columnName, mysql.TypeUnspecified, false, nil
	case *ast.AggregateFuncExpr:
		f, funcSql, t, err := parseAggregateFuncExpr(v, table)
		if err != nil {
			return "", "", 0, false, err
		}
		if !hasAsName {
			return "", "", 0, false, fmt.Errorf("aggregate function %q must have AS name", funcSql)
		}
		return f, funcSql, t, true, nil
	default:
		return "", "", mysql.TypeUnspecified, false, fmt.Errorf("unsupported select field: %T", v)
	}
}

func parseAggregateFuncExpr(node *ast.AggregateFuncExpr, from string) (string, string, byte, error) {
	funcName := node.F
	args := node.Args
	getColumnInfo := func() (string, string, error) {
		if len(args) == 0 {
			return "", "", fmt.Errorf("unsupported aggregate function: %s, missing args", funcName)
		}
		if len(args) > 1 {
			return "", "", fmt.Errorf("unsupported aggregate function: %s, expected one arg", funcName)
		}
		arg := args[0]
		switch v := arg.(type) {
		case *ast.ColumnNameExpr:
			columnName, err := parseColumn(v.Name, from)
			if err != nil {
				return "", "", err
			}

			return columnName, "", nil
		case *test_driver.ValueExpr:
			return "", fmt.Sprintf("%v", v.Datum.GetValue()), nil
		default:
			return "", "", nil
		}
	}

	name, value, err := getColumnInfo()
	if err != nil {
		return "", "", mysql.TypeUnspecified, err
	}

	var arg = name
	if len(value) > 0 {
		arg = value
	}
	var funcSql = fmt.Sprintf("%s(%s)", funcName, arg)
	tp, ok := funcMap[strings.ToLower(funcName)]
	if ok {
		return name, funcSql, tp, nil
	}

	return name, funcSql, mysql.TypeUnspecified, nil
}

func parseColumns(cols []*ast.ColumnName, table string) ([]string, error) {
	var columnSet = set.From()
	for _, col := range cols {
		colName, err := parseColumn(col, table)
		if err != nil {
			return nil, err
		}

		if colName != "" {
			columnSet.Add(colName)
		}
	}

	return columnSet.String(), nil
}

func parseColumn(col *ast.ColumnName, table string) (string, error) {
	if col == nil {
		return "", nil
	}
	colTable := col.Table.String()
	if len(colTable) > 0 && colTable != table {
		return "", fmt.Errorf("column table %q not match from %q", colTable, table)
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
