package spec

import "fmt"

func convertLimit(limit *Limit, table *Table, comment Comment) *Limit {
	if !limit.IsValid() {
		return limit
	}

	limit.TableInfo = table
	limit.Comment = comment
	return limit
}

func convertByItems(byItems []*ByItem, table *Table, comment Comment) ([]*ByItem, error) {
	var list []*ByItem
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
	column, ok := table.GetColumnByName(byItem.Column)
	if !ok {
		return nil, fmt.Errorf("column %q no found in table %q", byItem.Column, table.Name)
	}
	byItem.ColumnInfo = column
	return byItem, nil
}

func convertClause(clause *Clause, table *Table, comment Comment) (*Clause, error) {
	if !clause.IsValid() {
		return clause, nil
	}

	clause.Comment = comment
	clause.TableInfo = table
	column, ok := table.GetColumnByName(clause.Column)
	if !ok {
		return nil, fmt.Errorf("column %q no found in table %q", clause.Column, table.Name)
	}

	clause.ColumnInfo = column
	leftClause, err := convertClause(clause.Left, table, comment)
	if err != nil {
		return nil, err
	}
	rightClause, err := convertClause(clause.Right, table, comment)
	if err != nil {
		return nil, err
	}

	clause.Left = leftClause
	clause.Right = rightClause
	return clause, nil
}

func convertColumn(table *Table, columns []string) ([]Column, error) {
	var list []Column
	for _, c := range columns {
		column, ok := table.GetColumnByName(c)
		if !ok {
			return nil, fmt.Errorf("column %q no found in table %q", c, table.Name)
		}
		list = append(list, column)
	}
	return list, nil
}