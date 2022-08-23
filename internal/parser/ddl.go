package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseDDL(node *ast.CreateTableStmt) (*spec.DDL, error) {
	var ddl spec.DDL
	var err error
	ddl.Table, err = parseCreateTableStmt(node)
	if err != nil {
		return nil, err
	}
	return &ddl, nil
}
