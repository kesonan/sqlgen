package spec

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stringx"
)

// Table represents a table in the database.
type Table struct {
	// Columns is the list of columns in the table.
	Columns Columns
	// Constraint is a struct that contains the constraints of a table.
	// ConstraintForeignKey,ConstraintFulltext,ConstraintCheck are ignored.
	Constraint Constraint
	// Schema is the name of the schema that the table belongs to.
	Schema string
	// Name is the name of the table.
	Name string
}

type Columns []Column

// Column represents a column in a table.
type Column struct {
	// ColumnOption is a column option.
	ColumnOption
	// Name is the name of the column.
	Name string
	// TP is the type of the column.
	TP            byte
	AggregateCall bool
}

// ColumnOption is a column option.
type ColumnOption struct {
	// AutoIncrement is true if the column allows auto increment.
	AutoIncrement bool
	// Comment is the comment of the column.
	Comment string
	// HasDefault is true if the column has default value.
	HasDefaultValue bool
	// TODO: Add default value
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

// Has returns true if Columns has specified column.
func (cs Columns) Has(name string) bool {
	if cs == nil {
		return false
	}
	_, ok := cs.GetColumn(name)
	return ok
}

func (cs Columns) GetColumn(name string) (Column, bool) {
	for _, c := range cs {
		if c.Name == name {
			return c, true
		}
	}
	return Column{}, false
}

// IsPrimary returns true if the column is part of the primary key.
func (t *Table) IsPrimary(name string) bool {
	for _, c := range t.Constraint.PrimaryKey {
		if stringx.Contains(c, name) {
			return true
		}
	}
	return false
}

// ColumnList is a list of column names.
func (t *Table) ColumnList() []string {
	var list []string
	for _, c := range t.Columns {
		list = append(list, c.Name)
	}
	return list
}

// PrimaryColumnList is a list of column names that are part of the primary key.
func (t *Table) PrimaryColumnList() Columns {
	var ret Columns
	for _, list := range t.Constraint.PrimaryKey {
		for _, name := range list {
			c, ok := t.GetColumnByName(name)
			if !ok {
				continue
			}
			ret = append(ret, c)
		}
	}
	return ret
}

// PrimaryColumn returns the primary column.
func (t *Table) PrimaryColumn() Column {
	list := t.PrimaryColumnList()
	if len(list) == 0 {
		return Column{}
	}
	return list[0]
}

// HasOnePrimaryKey returns true if the table has one primary key.
func (t *Table) HasOnePrimaryKey() bool {
	return len(t.PrimaryColumnList()) == 1
}

// GetColumnByName returns the column with the given name.
func (t *Table) GetColumnByName(name string) (Column, bool) {
	for _, c := range t.Columns {
		if c.Name == name {
			return c, true
		}
	}
	return Column{}, false
}

func (t *Table) validate() error {
	if len(t.Name) == 0 {
		return fmt.Errorf("missing table name")
	}
	if len(t.Columns) == 0 {
		return fmt.Errorf("missing table columns")
	}
	if len(t.Constraint.PrimaryKey) == 0 {
		return fmt.Errorf("missing table primary key")
	}
	if len(t.Constraint.PrimaryKey) > 1 {
		return fmt.Errorf("unsupported multiple primary key")
	}
	return nil
}
