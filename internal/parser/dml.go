package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseDML(node ast.DMLNode) (spec.DML, error) {
	switch v := node.(type) {
	case *ast.InsertStmt:
		return parseInsert(v)
	case *ast.SelectStmt:
		return parseSelect(v)
	case *ast.DeleteStmt:
		return parseDelete(v)
	case *ast.UpdateStmt:
		return parseUpdate(v)
	default:
		return nil, errorUnsupportedStmt
	}
}

func parseTableRefsClause(clause *ast.TableRefsClause) (string, error) {
	if clause == nil {
		return "", errorMissingTable
	}

	var join = clause.TableRefs
	if join == nil {
		return "", errorMissingTable
	}

	if join.Left == nil {
		return "", errorMissingTable
	}

	if join.Right != nil {
		return "", errorMultipleTable
	}

	tableName, err := parseResultSetNode(join.Left)
	if err != nil {
		return "", err
	}

	return tableName, nil
}
