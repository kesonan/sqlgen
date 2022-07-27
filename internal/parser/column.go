package parser

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"

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
	column.TP = col.Tp
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
			column.Comment = opt.Expr.Text()
		case ast.ColumnOptionUniqKey:
			constraint.AppendUniqueKey(column.Name)
		case ast.ColumnOptionPrimaryKey:
			constraint.AppendPrimaryKey(column.Name)
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

	var keyName = constraint.Name
	var ret = spec.NewConstraint()
	switch constraint.Tp {
	case ast.ConstraintPrimaryKey:
		ret.AppendPrimaryKey(keyName, columns...)
	case ast.ConstraintKey, ast.ConstraintIndex:
		ret.AppendIndex(keyName, columns...)
	case ast.ConstraintUniq, ast.ConstraintUniqKey, ast.ConstraintUniqIndex:
		ret.AppendUniqueKey(keyName, columns...)
	default:
		// ignore other constraints
	}

	return &spec.Constraint{}
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
