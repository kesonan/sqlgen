package spec

import (
	"github.com/anqiansong/sqlgen/internal/set"
)

// NewConstraint returns a new Constraint.
func NewConstraint() *Constraint {
	return &Constraint{
		PrimaryKey: map[string][]string{},
		UniqueKey:  map[string][]string{},
		Index:      map[string][]string{},
	}
}

// AppendPrimaryKey appends a column to the primary key.
func (c *Constraint) AppendPrimaryKey(key string, columns ...string) {
	c.append(func(key string) ([]string, bool) {
		list, ok := c.PrimaryKey[key]
		return list, ok
	}, func(columns []string) {
		c.PrimaryKey[key] = columns
	}, key, columns...)
}

// AppendUniqueKey appends a column to the unique key.
func (c *Constraint) AppendUniqueKey(key string, columns ...string) {
	c.append(func(key string) ([]string, bool) {
		list, ok := c.UniqueKey[key]
		return list, ok
	}, func(columns []string) {
		c.UniqueKey[key] = columns
	}, key, columns...)
}

// AppendIndex appends a column to the unique key.
func (c *Constraint) AppendIndex(key string, columns ...string) {
	c.append(func(key string) ([]string, bool) {
		list, ok := c.Index[key]
		return list, ok
	}, func(columns []string) {
		c.Index[key] = columns
	}, key, columns...)
}

// IsEmpty returns true if the constraint is empty.
func (c *Constraint) IsEmpty() bool {
	return len(c.PrimaryKey) == 0 && len(c.UniqueKey) == 0 && len(c.Index) == 0
}

// Merge merges the constraint with another constraint.
func (c *Constraint) Merge(constraint *Constraint) {
	if constraint == nil {
		return
	}
	if constraint.IsEmpty() {
		return
	}

	for key, columns := range constraint.PrimaryKey {
		c.AppendPrimaryKey(key, columns...)
	}

	for key, columns := range constraint.UniqueKey {
		c.AppendUniqueKey(key, columns...)
	}

	for key, columns := range constraint.Index {
		c.AppendIndex(key, columns...)
	}
}

func (c *Constraint) append(existFn func(key string) ([]string, bool), result func(columns []string), key string, columns ...string) {
	var columnSet = set.FromString(columns...)
	if len(columns) == 0 {
		columns = []string{key}
	}

	list, ok := existFn(key)
	if !ok {
		result(columnSet.String())
		return
	}

	columnSet.AddStringList(list)
	for _, column := range columns {
		columnSet.Add(column)
	}

	result(columnSet.String())
}
