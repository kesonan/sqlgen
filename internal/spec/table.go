package spec

// Table represents a table in the database.
type Table struct {
	// Columns is the list of columns in the table.
	Columns []Column
	// Constraint is a struct that contains the constraints of a table.
	// ConstraintForeignKey,ConstraintFulltext,ConstraintCheck are ignored.
	Constraint Constraint
	// Name is the name of the table.
	Name string
}

// Column represents a column in a table.
type Column struct {
	// ColumnOption is a column option.
	ColumnOption
	// Name is the name of the column.
	Name string
	// TP is the type of the column.
	TP Type
}

// ColumnOption is a column option.
type ColumnOption struct {
	// AutoIncrement is true if the column allows auto increment.
	AutoIncrement bool
	// Comment is the comment of the column.
	Comment string
	// HasDefault is true if the column has default value.
	HasDefaultValue bool
	// NotNull is true if the column is not null, false represents the column is null.
	NotNull bool
	// Unsigned is true if the column is unsigned.
	Unsigned bool
}

// Constraint is a struct that contains the constraints of a table.
// ConstraintForeignKey,ConstraintFulltext,ConstraintCheck are ignored.
type Constraint struct {
	// Index is a list of column names that are part of an index, the key of map
	//	// is the key name, the values are the column list.
	Index map[string][]string
	// PrimaryKey is a list of column names that are part of the primary key, the key of map
	// is the key name, the values are the column list.
	PrimaryKey map[string][]string
	// UniqueKey is a list of column names that are part of a unique ke, the key of map
	//	// is the key name, the values are the column list.
	UniqueKey map[string][]string
}
