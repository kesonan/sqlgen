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

	// the below fields is convert from Statements
	Context
}

func (t Transaction) SQLText() string {
	return t.SQL
}

func (t Transaction) TableName() string {
	return ""
}

func (t Transaction) validate() (map[string]string, error) {
	return t.Context.validate()
}

func (t Transaction) HasArg() bool {
	for _, v := range t.InsertStmt {
		if v.HasArg() {
			return true
		}
	}
	for _, v := range t.SelectStmt {
		if v.HasArg() {
			return true
		}
	}
	for _, v := range t.UpdateStmt {
		if v.HasArg() {
			return true
		}
	}
	for _, v := range t.DeleteStmt {
		if v.HasArg() {
			return true
		}
	}
	return false
}
