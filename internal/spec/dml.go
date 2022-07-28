package spec

// DML represents a DML statement.
type DML interface {
	// SQLText returns the SQL text of the DML statement.
	SQLText() string
	validate() error
}
