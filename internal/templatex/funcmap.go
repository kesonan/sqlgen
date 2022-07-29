package templatex

import (
	"text/template"

	"github.com/iancoleman/strcase"
)

func UpperCamel(s string) string {
	return strcase.ToCamel(s)
}

func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

var funcMap = template.FuncMap{
	"upperCamel": UpperCamel,
	"lowerCamel": LowerCamel,
}
