package parser

import (
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/parser/test_driver"

	"github.com/anqiansong/sqlgen/internal/spec"
)

var p *parser.Parser

func init() {
	p = parser.New()
}

// Parse parses a SQL statement string and returns a spec.DXL.
func Parse(sql string) (*spec.DXL, error) {
	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	var ret spec.DXL
	for _, stmtNode := range stmtNodes {
		switch node := stmtNode.(type) {
		case ast.DDLNode:
			ddl, err := parseDDL(node)
			if err != nil {
				return nil, err
			}
			ret.DDL = append(ret.DDL, ddl)
		case ast.DMLNode:
			dml, err := parseDML(node)
			if err != nil {
				return nil, err
			}
			ret.DML = append(ret.DML, dml)
		default:
			// ignores other statements
		}
	}

	return &ret, nil
}
