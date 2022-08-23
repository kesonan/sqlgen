package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseDelete(stmt *ast.DeleteStmt, needFn bool) (spec.DML, error) {
	var ret spec.DeleteStmt
	var text = stmt.Text()
	comment, err := parseLineComment(text, needFn)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	sql, err := NewSqlScanner(text).ScanAndTrim()
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	if stmt.IsMultiTable {
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

	ret.Comment = comment
	ret.SQL = sql
	ret.Action = spec.ActionDelete
	ret.From = tableName
	return &ret, nil
}
