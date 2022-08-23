package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseUpdate(stmt *ast.UpdateStmt, needFn bool) (spec.DML, error) {
	var ret spec.UpdateStmt
	var text = stmt.Text()
	comment, err := parseLineComment(text, needFn)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	sql, err := NewSqlScanner(text).ScanAndTrim()
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	if stmt.MultipleTable {
		return nil, errorNearBy(errorMultipleTable, text)
	}

	tableName, err := parseTableRefsClause(stmt.TableRefs)
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

	if stmt.Order != nil {
		orderBy, err := parseOrderBy(stmt.Order, tableName)
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

	for _, a := range stmt.List {
		colName, err := parseColumn(a.Column, tableName)
		if err != nil {
			return nil, errorNearBy(err, text)
		}

		if len(colName) > 0 {
			ret.Columns = append(ret.Columns, colName)
		}
	}

	ret.Comment = comment
	ret.SQL = sql
	ret.Action = spec.ActionUpdate
	ret.Table = tableName

	return &ret, nil
}
