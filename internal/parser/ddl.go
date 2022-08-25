package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseDDL(node *ast.CreateTableStmt) (*spec.DDL, error) {
	var ddl spec.DDL
	ddl.Table = parseCreateTableStmt(node)
	return &ddl, nil
}
