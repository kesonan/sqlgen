package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseDDL(node ast.DDLNode) (*spec.DDL, error) {
	var ddl spec.DDL
	var err error
	switch stmt := node.(type) {
	case *ast.CreateTableStmt:
		ddl.Table, err = parseCreateTableStmt(stmt)
		if err != nil {
			return nil, err
		}
	default:
		// ignores other DDLs
	}
	return &ddl, nil
}
