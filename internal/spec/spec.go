package spec

// DXL describes ddl and dml.
type DXL struct {
	// DDL represents a DDL statement.
	DDL []*DDL
	// DML represents a DML statement.
	DML []DML
}

func (dxl *DXL) Stmt(table string) ([]*InsertStmt, []*UpdateStmt, []*SelectStmt, []*DeleteStmt) {
	var insertStmt []*InsertStmt
	var updateStmt []*UpdateStmt
	var selectStmt []*SelectStmt
	var deleteStmt []*DeleteStmt
	for _, dml := range dxl.DML {
		if dml.TableName() != table {
			continue
		}
		switch stmt := dml.(type) {
		case *InsertStmt:
			insertStmt = append(insertStmt, stmt)
		case *UpdateStmt:
			updateStmt = append(updateStmt, stmt)
		case *SelectStmt:
			selectStmt = append(selectStmt, stmt)
		case *DeleteStmt:
			deleteStmt = append(deleteStmt, stmt)
		}
	}

	return insertStmt, updateStmt, selectStmt, deleteStmt
}

// Validate validates the ddl and dml.
func (dxl *DXL) Validate() error {
	for _, ddl := range dxl.DDL {
		if err := ddl.validate(); err != nil {
			return err
		}
	}

	for _, dml := range dxl.DML {
		if err := dml.validate(); err != nil {
			return err
		}
	}
	return nil
}
