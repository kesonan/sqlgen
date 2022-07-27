package spec

// DXL describes ddl and dml.
type DXL struct {
	// DDL represents a DDL statement.
	DDL []*DDL
	// DML represents a DML statement.
	DML []DML
}
