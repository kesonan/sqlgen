package spec

// DXL describes ddl and dml.
type DXL struct {
	// DDL represents a DDL statement.
	DDL []*DDL
	// DML represents a DML statement.
	DML []DML
}

// Validate validates the ddl and dml.
func (dxl *DXL) Validate() error {
	for _, ddl := range dxl.DDL {
		if err := ddl.validate(); err != nil {
			return err
		}
	}

	for _, dml := range dxl.DML {
		if _, err := dml.validate(); err != nil {
			return err
		}
	}
	return nil
}
