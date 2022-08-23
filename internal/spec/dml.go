package spec

// DML represents a DML statement.
type DML interface {
	// SQLText returns the SQL text of the DML statement.
	SQLText() string
	// TableName returns the table of the DML statement.
	TableName() string
	validate() (map[string]string, error)
}
