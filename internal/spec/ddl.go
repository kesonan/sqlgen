package spec

// DDL represents a DDL statement.
type DDL struct {
	// Table represents a table in the database.
	Table *Table
}

// IsEmpty returns true if the DDL is empty.
func (d *DDL) IsEmpty() bool {
	return d.Table == nil
}

func (d *DDL) validate() error {
	if d.Table == nil {
		return nil
	}
	return d.Table.validate()
}
