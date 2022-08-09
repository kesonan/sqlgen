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
var sqlScannerTpl string

func Run(list []spec.Context, output string) error {
	for _, ctx := range list {
		var genFilename = filepath.Join(output, fmt.Sprintf("%s_model.gen.go", ctx.Table.Name))
		var customFilename = filepath.Join(output, fmt.Sprintf("%s_model.go", ctx.Table.Name))
		var scannerFilename = filepath.Join(output, "scanner.go")
		gen := templatex.New()
		var columns, parameter, args []string
		for _, c := range ctx.Table.Columns {
			if c.AutoIncrement {
				continue
			}
			columns = append(columns, fmt.Sprintf("`%s`", c.Name))
			parameter = append(parameter, "?")
			args = append(args, strcase.ToCamel(c.Name))
		}
		gen.AppendFuncMap(template.FuncMap{
			"IsPrimary": func(name string) bool {
				return ctx.Table.IsPrimary(name)
			},
			"InsertSQL": func() string {
				return fmt.Sprintf(`"INSERT INTO %s (%s) VALUES (%s)"`, fmt.Sprintf("`%s`", ctx.Table.Name), strings.Join(columns, ", "), strings.Join(parameter, ", "))
			},
			"InsertSQLArgs": func(pkg string) string {
				return pkg + "." + strings.Join(args, fmt.Sprintf(", %s.", pkg))
			},
			"HavingSprintf": func(format string) string {
				format = strings.ReplaceAll(format, "?", "%v")
				return format
			},
		})
		scanner := templatex.New()
		scanner.MustParse(sqlScannerTpl)
		scanner.MustExecute(ctx)
		scanner.MustSave(scannerFilename, true)

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
