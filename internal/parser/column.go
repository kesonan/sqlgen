package parser

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/test_driver"

	"github.com/anqiansong/sqlgen/internal/set"
	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseColumnDef(col *ast.ColumnDef) (*spec.Column, *spec.Constraint) {
	if col == nil || col.Name == nil {
		return nil, nil
	}

	var column spec.Column
	var constraint = spec.NewConstraint()
	var tp = col.Tp
	if tp != nil {
		column.Unsigned = mysql.HasUnsignedFlag(tp.Flag)
		column.TP = tp.Tp
	}

	column.Name = col.Name.String()
	for _, opt := range col.Options {
		var tp = opt.Tp
		switch tp {
		case ast.ColumnOptionNotNull:
			column.NotNull = true
		case ast.ColumnOptionAutoIncrement:
			column.AutoIncrement = true
		case ast.ColumnOptionDefaultValue:
			column.HasDefaultValue = true
		case ast.ColumnOptionComment:
			var expr = opt.Expr
			if expr != nil {
				value, ok := expr.(*test_driver.ValueExpr)
				if ok {
					column.Comment = value.GetString()
				}
			}
		case ast.ColumnOptionUniqKey:
			constraint.AppendUniqueKey(column.Name, column.Name)
		case ast.ColumnOptionPrimaryKey:
			constraint.AppendPrimaryKey(column.Name, column.Name)
		default:
			// ignore other options
		}
	}

	return &column, constraint
}

func parseConstraint(constraint *ast.Constraint) *spec.Constraint {
	if constraint == nil {
		return nil
	}

	var columns = parseColumnFromKeys(constraint.Keys)
	if len(columns) == 0 {
		return nil
	}

	var ret = spec.NewConstraint()
	var key = constraint.Name
	switch constraint.Tp {
	case ast.ConstraintPrimaryKey:
		ret.AppendPrimaryKey(key, columns...)
	case ast.ConstraintKey, ast.ConstraintIndex:
		ret.AppendIndex(key, columns...)
	case ast.ConstraintUniq, ast.ConstraintUniqKey, ast.ConstraintUniqIndex:
		ret.AppendUniqueKey(key, columns...)
	default:
		// ignore other constraints
	}

	return ret
}

func parseColumnFromKeys(keys []*ast.IndexPartSpecification) []string {
	var columnSet = set.From()
	for _, key := range keys {
		if key.Column == nil {
			continue
		}

		var columnName = key.Column.String()
		columnSet.Add(columnName)
	}

	return columnSet.String()
}
