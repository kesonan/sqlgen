package spec

import (
	"fmt"

	"github.com/pingcap/parser/mysql"
)

func convertLimit(limit *Limit, table *Table, comment Comment) *Limit {
	if !limit.IsValid() {
		return limit
	}

	limit.TableInfo = table
	limit.Comment = comment
	return limit
}

func convertByItems(byItems ByItems, table *Table, comment Comment) (ByItems, error) {
	var list ByItems
	for _, v := range byItems {
		byItem, err := convertByItem(v, table, comment)
		if err != nil {
			return nil, err
		}
		list = append(list, byItem)
	}
	return list, nil
}

func convertByItem(byItem *ByItem, table *Table, comment Comment) (*ByItem, error) {
	if !byItem.IsValid() {
		return byItem, nil
	}

	byItem.TableInfo = table
	byItem.Comment = comment
	if byItem.Column == WildCard {
		return nil, fmt.Errorf("wildcard is not allowed in by item")
	}
	column, ok := table.GetColumnByName(byItem.Column)
	if !ok {
		return nil, fmt.Errorf("column %q no found in table %q", byItem.Column, table.Name)
	}
	byItem.ColumnInfo = column
	return byItem, nil
}

func convertClause(clause *Clause, table *Table, comment Comment, rows Columns) (*Clause, error) {
	if !clause.IsValid() {
		return clause, nil
	}

	clause.Comment = comment
	clause.TableInfo = table
	if clause.Column == WildCard {
		return nil, fmt.Errorf("wildcard is not allowed in by item")
	}
	if len(clause.Column) > 0 {
		column, ok := table.GetColumnByName(clause.Column)
		if ok {
			clause.ColumnInfo = column
		} else {
			// for case: select max(id) AS maxID from t having maxID > 0;
			column, ok = rows.GetColumn(clause.Column)
			if !ok {
				return nil, fmt.Errorf("column %q no found in table %q", clause.Column, table.Name)
			}
		}

		clause.ColumnInfo = column
	}

	leftClause, err := convertClause(clause.Left, table, comment, rows)
	if err != nil {
		return nil, err
	}
	rightClause, err := convertClause(clause.Right, table, comment, rows)
	if err != nil {
		return nil, err
	}

	clause.Left = leftClause
	clause.Right = rightClause
	return clause, nil
}

func convertColumn(table *Table, columns []string) Columns {
	var list Columns
	var m = map[string]struct{}{}
	for _, c := range columns {
		if _, ok := m[c]; ok {
			continue
		}
		if c == WildCard {
			list = append(list, table.Columns...)
			continue
		}

		column, ok := table.GetColumnByName(c)
		if ok {
			list = append(list, column)
		}

	}
	return list
}

func convertField(table *Table, fields []Field) (Columns, error) {
	var list Columns
	var m = map[string]struct{}{}
	for _, f := range fields {
		name := f.ColumnName
		if len(f.ASName) > 0 {
			name = f.ASName
		}
		if _, ok := m[name]; ok {
			continue
		}
		m[name] = struct{}{}
		if name == WildCard {
			list = append(list, table.Columns...)
			continue
		}

		if len(f.ColumnName) > 0 {
			column, ok := table.GetColumnByName(f.ColumnName)
			if ok {
				column.Name = name
				column.AggregateCall = f.AggregateCall
				if f.TP != mysql.TypeUnspecified {
					column.TP = f.TP
				}
				list = append(list, column)
			} else {
				return nil, fmt.Errorf("column %q no found in table %q", f.ColumnName, table.Name)
			}
		} else {
			if f.TP == mysql.TypeUnspecified {
				return nil, fmt.Errorf("column %q no found in table %q", f.ColumnName, table.Name)
			}
			list = append(list, Column{
				Name:          name,
				TP:            f.TP,
				AggregateCall: f.AggregateCall,
			})
		}
	}
	return list, nil
}
