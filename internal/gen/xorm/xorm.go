package xorm

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/anqiansong/sqlgen/internal/templatex"
)

//go:embed xorm_gen.tpl
var xormGenTpl string

//go:embed xorm_custom.tpl
var xormCustomTpl string

func Run(list []spec.Context, output string) error {
	for _, ctx := range list {
		var genFilename = filepath.Join(output, fmt.Sprintf("%s_model.gen.go", ctx.Table.Name))
		var customFilename = filepath.Join(output, fmt.Sprintf("%s_model.go", ctx.Table.Name))
		gen := templatex.New()
		gen.AppendFuncMap(template.FuncMap{
			"IsPrimary": func(name string) bool {
				return ctx.Table.IsPrimary(name)
			},
			"HavingSprintf": func(format string) string {
				format = strings.ReplaceAll(format, "?", "'%v'")
				return format
			},
			"IsExtraResult": func(name string) bool {
				return name != templatex.UpperCamel(ctx.Table.Name)
			},
		})
		gen.MustParse(xormGenTpl)
		gen.MustExecute(ctx)
		gen.MustSaveAs(genFilename, true)

		custom := templatex.New()
		custom.MustParse(xormCustomTpl)
		custom.MustExecute(ctx)
		custom.MustSave(customFilename, true)
	}

	return nil
}
