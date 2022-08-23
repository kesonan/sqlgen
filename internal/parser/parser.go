package parser

import (
	"fmt"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/parser/test_driver"

	"github.com/anqiansong/sqlgen/internal/spec"
)

var p *parser.Parser

type stmts []stmt

type stmt interface {
	nodes() []ast.StmtNode
}

type createTableStmt struct {
	stmt *ast.CreateTableStmt
}

func (c createTableStmt) nodes() []ast.StmtNode {
	return []ast.StmtNode{c.stmt}
}

type queryStmt struct {
	stmt ast.StmtNode
}

func (q queryStmt) nodes() []ast.StmtNode {
	return []ast.StmtNode{q.stmt}
}

type transactionStmt struct {
	startTransactionStmt ast.StmtNode
	queryList            stmts
	commitStmt           ast.StmtNode
}

func (t transactionStmt) nodes() []ast.StmtNode {
	var list []ast.StmtNode
	for _, v := range t.queryList {
		stmt, ok := v.(*queryStmt)
		if ok {
			list = append(list, stmt.stmt)
		}
	}
	return list
}

func init() {
	p = parser.New()
}

// Parse parses a SQL statement string and returns a spec.DXL.
func Parse(sql string) (*spec.DXL, error) {
	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	stmt, err := splits(stmtNodes)
	if err != nil {
		return nil, err
	}

	var ret spec.DXL
	for _, stmtNode := range stmt {
		switch node := stmtNode.(type) {
		case *createTableStmt:
			ddl, err := parseDDL(node.stmt)
			if err != nil {
				return nil, err
			}
			ret.DDL = append(ret.DDL, ddl)
		case *queryStmt:
			dml, err := parseDML(node.stmt, true)
			if err != nil {
				return nil, err
			}
			ret.DML = append(ret.DML, dml)
		case *transactionStmt:
			if node.queryList.hasTransactionStmt() {
				return nil, errorUnsupportedNestedTransaction
			}
			if len(node.queryList) == 0 {
				continue
			}
			dml, err := parseTransaction(node)
			if err != nil {
				return nil, err
			}
			ret.DML = append(ret.DML, dml)
		default:
			// ignores other statements
		}
	}

	if err = ret.Validate(); err != nil {
		return nil, err
	}

	return &ret, nil
}

func splits(stmtNodes []ast.StmtNode) ([]stmt, error) {
	var list stmts
	var transactionMode bool
	for _, v := range stmtNodes {
		switch node := v.(type) {
		case *ast.CreateTableStmt:
			if transactionMode {
				return nil, fmt.Errorf("missing begin stmt near by '%s'", v.Text())
			}
			list = append(list, &createTableStmt{stmt: node})
		case *ast.InsertStmt, *ast.SelectStmt, *ast.DeleteStmt, *ast.UpdateStmt:
			if transactionMode {
				transactionNode := list[len(list)-1].(*transactionStmt)
				transactionNode.queryList = append(transactionNode.queryList, &queryStmt{stmt: node})
			} else {
				list = append(list, &queryStmt{stmt: node})
			}
		case *ast.BeginStmt:
			if transactionMode {
				transactionNode := list[len(list)-1].(*transactionStmt)
				transactionNode.queryList = append(transactionNode.queryList, &transactionStmt{startTransactionStmt: node})
			} else {
				transactionMode = true
				list = append(list, &transactionStmt{startTransactionStmt: v})
			}
		case *ast.CommitStmt:
			if transactionMode {
				transactionNode := list[len(list)-1].(*transactionStmt)
				transactionNode.commitStmt = v
				transactionMode = false
			} else {
				return nil, fmt.Errorf("missing begin stmt near by '%s'", v.Text())
			}
		default:
			return nil, errorUnsupportedStmt
		}
	}
	if transactionMode {
		return nil, errorMissingCommit
	}
	return list, nil
}

func (s stmts) hasTransactionStmt() bool {
	for _, v := range s {
		if _, ok := v.(*transactionStmt); ok {
			return true
		}
	}
	return false
}
