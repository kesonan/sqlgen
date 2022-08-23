package spec

type Transaction struct {
	// Action represents the db action.
	Action Action
	// Comment represents a sql comment.
	Comment
	// SQL represents the original sql text.
	SQL string
	// Statements represents the list of statement.
	Statements []DML
}

func (t Transaction) SQLText() string {
	return t.SQL
}

func (t Transaction) TableName() string {
	return ""
}
