package spec

import "fmt"

// Context is sql table and query context.
type Context struct {
	Table       *Table
	InsertStmt  []*InsertStmt
	SelectStmt  []*SelectStmt
	UpdateStmt  []*UpdateStmt
	DeleteStmt  []*DeleteStmt
	Transaction []*Transaction
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
		transaction, isTransaction := d.(*Transaction)
		if !isTransaction {
			if d.TableName() == table.Name {
				switch v := d.(type) {
				case *InsertStmt:
					columns := convertColumn(table, v.Columns)
					v.ColumnInfo = columns
					v.TableInfo = table
					ctx.InsertStmt = append(ctx.InsertStmt, v)
				case *SelectStmt:
					columns, err := convertField(table, v.Columns)
					if err != nil {
						return ctx, err
					}

					v.GroupBy, err = convertByItems(v.GroupBy, table, v.Comment)
					if err != nil {
						return ctx, err
					}

					v.Having, err = convertClause(v.Having, table, v.Comment, columns)
					if err != nil {
						return ctx, err
					}

					v.Where, err = convertClause(v.Where, table, v.Comment, columns)
					if err != nil {
						return ctx, err
					}

					v.OrderBy, err = convertByItems(v.OrderBy, table, v.Comment)
					if err != nil {
						return ctx, err
					}

					v.Limit = convertLimit(v.Limit, table, v.Comment)
					v.ColumnInfo = columns
					v.FromInfo = table
					ctx.SelectStmt = append(ctx.SelectStmt, v)
				case *UpdateStmt:
					columns := convertColumn(table, v.Columns)
					var err error

					v.Where, err = convertClause(v.Where, table, v.Comment, columns)
					if err != nil {
						return ctx, err
					}

					v.OrderBy, err = convertByItems(v.OrderBy, table, v.Comment)
					if err != nil {
						return ctx, err
					}

					v.TableInfo = table
					v.ColumnInfo = columns
					v.Limit = convertLimit(v.Limit, table, v.Comment)
					ctx.UpdateStmt = append(ctx.UpdateStmt, v)
				case *DeleteStmt:
					var err error
					v.Where, err = convertClause(v.Where, table, v.Comment, nil)
					if err != nil {
						return ctx, err
					}

					v.OrderBy, err = convertByItems(v.OrderBy, table, v.Comment)
					if err != nil {
						return ctx, err
					}

					v.FromInfo = table
					v.Limit = convertLimit(v.Limit, table, v.Comment)
					ctx.DeleteStmt = append(ctx.DeleteStmt, v)
				}
			}
		} else {
			childCtx, err := from(table, transaction.Statements)
			if err != nil {

				return Context{}, err
			}
			transaction.Context = childCtx
			ctx.Transaction = append(ctx.Transaction, transaction)
		}
	}
	if _, err := ctx.validate(); err != nil {
		return Context{}, err
	}
	return ctx, nil
}

func (ctx Context) validate() (map[string]string, error) {
	funcName := map[string]string{}
	for _, v := range ctx.InsertStmt {
		if _, err := v.validate(); err != nil {
			return nil, err
		}
		if _, ok := funcName[v.FuncName]; ok {
			return nil, fmt.Errorf("duplicate function %q near by %q", v.FuncName, v.OriginText)
		}
		funcName[v.FuncName] = v.OriginText
	}
	for _, v := range ctx.SelectStmt {
		if _, err := v.validate(); err != nil {
			return nil, err
		}
		if _, ok := funcName[v.FuncName]; ok {
			return nil, fmt.Errorf("duplicate function %q near by %q", v.FuncName, v.OriginText)
		}
		funcName[v.FuncName] = v.OriginText
	}
	for _, v := range ctx.UpdateStmt {
		if _, err := v.validate(); err != nil {
			return nil, err
		}
		if _, ok := funcName[v.FuncName]; ok {
			return nil, fmt.Errorf("duplicate function %q near by %q", v.FuncName, v.OriginText)
		}
		funcName[v.FuncName] = v.OriginText
	}
	for _, v := range ctx.DeleteStmt {
		if _, err := v.validate(); err != nil {
			return nil, err
		}
		if _, ok := funcName[v.FuncName]; ok {
			return nil, fmt.Errorf("duplicate function %q near by %q", v.FuncName, v.OriginText)
		}
		funcName[v.FuncName] = v.OriginText
	}
	for _, v := range ctx.Transaction {
		if funcM, err := v.validate(); err != nil {
			return nil, err
		} else {
			for fn, originText := range funcM {
				if _, ok := funcName[fn]; ok {
					return nil, fmt.Errorf("duplicate function %q near by %q", fn, originText)
				}
				funcName[fn] = originText
			}
		}
	}
	return funcName, nil
}
