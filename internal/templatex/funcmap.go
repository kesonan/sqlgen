package templatex

import (
	"strings"
	"text/template"

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

var funcMap = template.FuncMap{
	"UpperCamel": UpperCamel,
	"LowerCamel": LowerCamel,
	"Join":       Join,
}
