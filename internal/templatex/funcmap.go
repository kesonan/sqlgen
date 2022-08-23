package templatex

import (
	"strings"
	"text/template"

	"github.com/anqiansong/sqlgen/internal/stringx"
	"github.com/iancoleman/strcase"
)

func UpperCamel(s string) string {
	return strcase.ToCamel(s)
}

func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

func Join(list []string, sep string) string {
	return strings.Join(list, sep)
}

func LineComment(s string) string {
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n'
	})
	return strings.Join(fields, "\n// ")
}

var funcMap = template.FuncMap{
	"UpperCamel":  UpperCamel,
	"LowerCamel":  LowerCamel,
	"Join":        Join,
	"TrimNewLine": stringx.TrimNewLine,
	"LineComment": LineComment,
}
