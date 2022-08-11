package spec

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/anqiansong/sqlgen/internal/buffer"
	"github.com/anqiansong/sqlgen/internal/templatex"
)

// SelectStmt represents a select statement.
type SelectStmt struct {
	// Action represents the db action.
	Action Action
	// SelectSQL represents the select filed sql.
	SelectSQL string
	// Columns represents the operation columns.
	Columns Fields
	// Comment represents a sql comment.
	Comment
	// Distinct represents the select distinct flag.
	Distinct bool
	// From represents the operation table name, do not support multiple tables.
	From string
	// GroupBy represents the group by clause.
	GroupBy ByItems
	// Having represents the having clause.
	Having *Clause
	// Limit represents the limit clause.
	Limit *Limit
	// OrderBy represents the order by clause.
	OrderBy ByItems
	// SQL represents the original sql text.
	SQL string
	// Where represents the where clause.
	Where *Clause

	// the below data are from table
	// ColumnInfo are the column info which are convert from Columns.
	ColumnInfo Columns
	// FromInfo is the table info which is convert from From.
	FromInfo *Table
}

func (s *SelectStmt) SQLText() string {
	return s.SQL
}

func (s *SelectStmt) TableName() string {
	return s.From
}

func (s *SelectStmt) ReceiverName() string {
	if s.ContainsExtraColumns() {
		return strcase.ToCamel(fmt.Sprintf("%sResult", s.FuncName))
	}
	return strcase.ToCamel(s.TableName())
}

//go:embed column.tpl
var fieldTpl string

func (s *SelectStmt) ReceiverStructure(orm string) string {
	receiverName := s.ReceiverName()
	if strings.EqualFold(receiverName, s.TableName()) {
		// Use table struct
		return ""
	}
	var buf = buffer.New()
	buf.Write("\n")
	buf.Write("// %s is a %s.", receiverName, strcase.ToDelimited(receiverName, ' '))
	buf.Write(`type %s struct {`, receiverName)
	for _, v := range s.ColumnInfo {
		t := templatex.New()
		t.AppendFuncMap(template.FuncMap{
			"ColumnTag": func() string {
				switch orm {
				case "gorm":
					return fmt.Sprintf(`gorm:"column:%s" `, v.Name)
				case "xorm":
					return fmt.Sprintf(`xorm:"'%s'" `, v.Name)
				default:
					return ""
				}
			},
		})
		t.MustParse(fieldTpl)
		t.MustExecute(v)
		var columnBuf bytes.Buffer
		t.Write(&columnBuf, false)
		buf.Write(columnBuf.String())
	}
	buf.Write("}")
	return buf.String()
}

// ContainsExtraColumns returns true if the select statement contains extra columns.
func (s *SelectStmt) ContainsExtraColumns() bool {
	for _, f := range s.Columns {
		name := f.Name()
		if name == WildCard {
			continue
		}
		if !s.FromInfo.Columns.Has(name) {
			return true
		}
	}
	return false
}

func (s *SelectStmt) validate() error {
	return s.Comment.validate()
}
