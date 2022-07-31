package spec

// Context is sql table and query context.
type Context struct {
	Table      *Table
	InsertStmt []*InsertStmt
	SelectStmt []*SelectStmt
	UpdateStmt []*UpdateStmt
	DeleteStmt []*DeleteStmt
}

// From creates context from table and dml.
func From(dxl *DXL) ([]Context, error) {
	var list []Context
	for _, d := range dxl.DDL {
		ctx, err := from(d.Table, dxl.DML)
		if err != nil {
			return nil, err
		}

		list = append(list, ctx)
	}

	return list, nil
}

func from(table *Table, dml []DML) (Context, error) {
	var ctx Context
	ctx.Table = table
	for _, d := range dml {
		if d.TableName() == table.Name {
			switch v := d.(type) {
			case *InsertStmt:
				columns, err := convertColumn(table, v.Columns)
				if err != nil {
					return Context{}, err
				}
				v.ColumnInfo = columns
				v.TableInfo = table
				ctx.InsertStmt = append(ctx.InsertStmt, v)
			case *SelectStmt:
				columns, err := convertColumn(table, v.Columns)
				if err != nil {
					return Context{}, err
				}

				groupByColumns, err := convertColumn(table, v.GroupBy)
				if err != nil {
					return Context{}, err
				}

				v.Having, err = convertClause(v.Having, table)
				if err != nil {
					return Context{}, err
				}

				v.Where, err = convertClause(v.Where, table)
				if err != nil {
					return Context{}, err
				}

				v.OrderBy, err = convertByItems(v.OrderBy, table)
				if err != nil {
					return Context{}, err
				}

				v.Limit = convertLimit(v.Limit, table)
				v.ColumnInfo = columns
				v.GroupByInfo = groupByColumns
				v.FromInfo = table
				ctx.SelectStmt = append(ctx.SelectStmt, v)
			case *UpdateStmt:
				columns, err := convertColumn(table, v.Columns)
				if err != nil {
					return Context{}, err
				}

				v.Where, err = convertClause(v.Where, table)
				if err != nil {
					return Context{}, err
				}

				v.OrderBy, err = convertByItems(v.OrderBy, table)
				if err != nil {
					return Context{}, err
				}

				v.TableInfo = table
				v.ColumnInfo = columns
				v.Limit = convertLimit(v.Limit, table)
				ctx.UpdateStmt = append(ctx.UpdateStmt, v)
			case *DeleteStmt:
				var err error
				v.OrderBy, err = convertByItems(v.OrderBy, table)
				if err != nil {
					return Context{}, err
				}

				v.FromInfo = table
				v.Limit = convertLimit(v.Limit, table)
				ctx.DeleteStmt = append(ctx.DeleteStmt, v)
			}
		}
	}
	return ctx, nil
}
