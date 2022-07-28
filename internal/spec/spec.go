package spec

// DXL describes ddl and dml.
type DXL struct {
	// DDL represents a DDL statement.
	DDL []*DDL
	// DML represents a DML statement.
	DML []DML
}

// Validate validates the ddl and dml.
func (xml *DXL) Validate() error {
	for _, ddl := range xml.DDL {
		if err := ddl.validate(); err != nil {
			return err
		}
	}

	for _, dml := range xml.DML {
		if err := dml.validate(); err != nil {
			return err
		}
	}
	return nil
}
