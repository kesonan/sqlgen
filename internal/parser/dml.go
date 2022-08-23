package parser

import (
	"github.com/anqiansong/sqlgen/internal/buffer"
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseDML(node ast.StmtNode) (spec.DML, error) {
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

func parseTransaction(node *transactionStmt) (spec.DML, error) {
	if node == nil {
		return nil, errorMissingTransaction
	}
	var sqlBuilder = buffer.New()
	var beginText = node.startTransactionStmt.Text()
	var commitText = node.commitStmt.Text()
	beginSQL, err := NewSqlScanner(beginText).ScanAndTrim()
	if err != nil {
		return nil, errorNearBy(err, beginText)
	}
	commitSQL, err := NewSqlScanner(commitText).ScanAndTrim()
	if err != nil {
		return nil, errorNearBy(err, commitText)
	}

	comment, err := parseLineComment(beginText)
	if err != nil {
		return nil, err
	}

	sqlBuilder.Write(beginSQL)
	var ret spec.Transaction
	ret.Action = spec.ActionTransaction
	for _, v := range node.nodes() {
		dml, err := parseDML(v)
		if err != nil {
			return nil, err
		}
		sqlBuilder.Write(dml.SQLText())
		ret.Statements = append(ret.Statements, dml)
	}
	sqlBuilder.Write(commitSQL)
	ret.SQL = sqlBuilder.String()
	ret.Comment = comment
	return &ret, nil
}
