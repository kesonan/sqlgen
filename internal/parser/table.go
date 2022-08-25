package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseCreateTableStmt(stmt *ast.CreateTableStmt) *spec.Table {
	var table spec.Table
	if stmt.Table != nil {
		table.Name = stmt.Table.Name.String()
	}

	var constraint = spec.NewConstraint()
	for _, col := range stmt.Cols {
		column, con := parseColumnDef(col)
		if column != nil {
			table.Columns = append(table.Columns, *column)
		}
		constraint.Merge(con)
	}

	for _, c := range stmt.Constraints {
		constraint.Merge(parseConstraint(c))
	}

	table.Constraint = *constraint
	return &table
}
