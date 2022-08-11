package sql

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/anqiansong/sqlgen/internal/templatex"
)

//go:embed sql_gen.tpl
var sqlGenTpl string

//go:embed sql_custom.tpl
var sqlCustomTpl string

//go:embed scanner.tpl
var scannerTpl string

func Run(list []spec.Context, output string) error {
	var scannerFilename = filepath.Join(output, "scanner.go")
	scanner := templatex.New()
	scanner.MustParse(scannerTpl)
	scanner.MustExecute(nil)
	scanner.MustSave(scannerFilename, true)

	for _, ctx := range list {
		var genFilename = filepath.Join(output, fmt.Sprintf("%s_model.gen.go", ctx.Table.Name))
		var customFilename = filepath.Join(output, fmt.Sprintf("%s_model.go", ctx.Table.Name))
		gen := templatex.New()
		var insertQuery, insertQuotes []string
		for _, v := range ctx.Table.Columns {
			if v.AutoIncrement {
				continue
			}
			insertQuery = append(insertQuery, fmt.Sprintf("`%s`", v.Name))
			insertQuotes = append(insertQuotes, "?")
		}
		gen.AppendFuncMap(template.FuncMap{
			"IsPrimary": func(name string) bool {
				return ctx.Table.IsPrimary(name)
			},
			"InsertSQL": func() string {
				return strings.Join(insertQuery, ", ")
			},
			"InsertQuotes": func() string {
				return strings.Join(insertQuotes, ", ")
			},
			"InsertValues": func(pkg string) string {
				var values []string
				for _, v := range ctx.Table.Columns {
					if v.AutoIncrement {
						continue
					}
					values = append(values, fmt.Sprintf("%s.%s", pkg, strcase.ToCamel(v.Name)))
				}
				return strings.Join(values, ", ")
			},
			"HavingSprintf": func(format string) string {
				format = strings.ReplaceAll(format, "?", "%v")
				return format
			},
		})
		gen.MustParse(sqlGenTpl)
		gen.MustExecute(ctx)
		gen.MustSaveAs(genFilename, true)

		custom := templatex.New()
		custom.MustParse(sqlCustomTpl)
		custom.MustExecute(ctx)
		custom.MustSave(customFilename, true)
	}
	return nil
}
