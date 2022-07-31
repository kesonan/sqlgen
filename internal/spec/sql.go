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
				ctx.InsertStmt = append(ctx.InsertStmt, v)
			case *SelectStmt:
				ctx.SelectStmt = append(ctx.SelectStmt, v)
			case *UpdateStmt:
				ctx.UpdateStmt = append(ctx.UpdateStmt, v)
			case *DeleteStmt:
				ctx.DeleteStmt = append(ctx.DeleteStmt, v)
			}
		}
	}
	return ctx, nil
}
