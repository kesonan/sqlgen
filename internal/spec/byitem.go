package spec

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/anqiansong/sqlgen/internal/buffer"
	"github.com/anqiansong/sqlgen/internal/parameter"
	"github.com/anqiansong/sqlgen/internal/set"
)

// ByItems returns the by items.
type ByItems []*ByItem

// ByItem represents an order-by or group-by item.
type ByItem struct {
	// Column represents the column name.
	Column string
	// Desc returns true if order by Column desc.
	Desc bool

	// the below data are from table
	// ColumnInfo are the column info which are convert from Column.
	ColumnInfo Column
	// TableInfo is the table info.
	TableInfo *Table

	// the below data are from stmt
	// Comment represents a sql comment.
	Comment Comment
}

func (b *ByItem) IsValid() bool {
	if b == nil {
		return false
	}

	return len(b.Column) > 0
}

func (b ByItems) IsValid() bool {
	if len(b) == 0 {
		return false
	}

	for _, v := range b {
		if v.IsValid() {
			return true
		}
	}

	return false
}

// SQL returns the clause condition strings.
func (b ByItems) SQL() (string, error) {
	sql, _, err := b.marshal()
	return fmt.Sprintf("`%s`", sql), err
}

// ParameterStructure returns the parameter type structure.
func (b ByItems) ParameterStructure(identifier string) (string, error) {
	_, parameters, err := b.marshal()
	if err != nil {
		return "", err
	}

	var writer = buffer.New()
	writer.Write(`// %s is a %s parameter structure.`, b.ParameterStructureName(identifier), strcase.ToDelimited(identifier, ' '))
	writer.Write(`type %s struct {`, b.ParameterStructureName(identifier))
	for _, v := range parameters {
		writer.Write("%s %s", v.Column, v.Type)
	}

	writer.Write(`}`)

	return writer.String(), nil
}

// ParameterThirdImports returns the third package imports.
func (b ByItems) ParameterThirdImports() (string, error) {
	_, parameters, err := b.marshal()
	if err != nil {
		return "", err
	}
	var thirdPkgSet = set.From()
	for _, v := range parameters {
		if len(v.ThirdPkg) == 0 {
			continue
		}
		thirdPkgSet.Add(v.ThirdPkg)
	}

	return strings.Join(thirdPkgSet.String(), "\n"), nil
}

// Parameters returns the parameter variables.
func (b ByItems) Parameters(pkg string) (string, error) {
	_, parameters, err := b.marshal()
	if err != nil {
		return "", err
	}
	var list []string
	for _, v := range parameters {
		list = append(list, fmt.Sprintf("%s.%s", pkg, v.Column))
	}

	return strings.Join(list, ", "), nil
}

// ParameterStructureName returns the parameter structure name.
func (b ByItems) ParameterStructureName(identifier string) string {
	if !b.IsValid() {
		return ""
	}

	one := b[0]
	return strcase.ToCamel(fmt.Sprintf("%s%sParameter", one.Comment.FuncName, identifier))
}

func (b ByItems) marshal() (sql string, parameters parameter.Parameters, err error) {
	parameters = parameter.Empty
	if len(b) == 0 {
		return
	}

	var sqlJoin []string
	var ps = parameter.New()
	for _, v := range b {
		if v.Desc {
			sqlJoin = append(sqlJoin, fmt.Sprintf("%s desc", v.Column))
		} else {
			sqlJoin = append(sqlJoin, v.Column)
		}

		p, err := v.ColumnInfo.DataType()
		if err != nil {
			return "", nil, err
		}

		ps.Add(p)
	}

	sql = strings.Join(sqlJoin, ", ")
	parameters = ps.List()
	return
}
