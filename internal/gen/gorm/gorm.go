package gorm

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/anqiansong/sqlgen/internal/templatex"
)

//go:embed gorm_gen.tpl
var gormGenTpl string

//go:embed gorm_custom.tpl
var gormCustomTpl string

func Run(list []spec.Context, output string) error {
	for _, ctx := range list {
		var genFilename = filepath.Join(output, fmt.Sprintf("%s_model.gen.go", ctx.Table.Name))
		var customFilename = filepath.Join(output, fmt.Sprintf("%s_model.go", ctx.Table.Name))
		gen := templatex.New()
		gen.AppendFuncMap(funcMap)
		gen.AppendFuncMap(template.FuncMap{
			"IsPrimary": func(name string) bool {
				return ctx.Table.IsPrimary(name)
			},
		})
		gen.MustParse(gormGenTpl)
		gen.MustExecute(ctx)
		gen.MustSaveAs(genFilename, true)

		custom := templatex.New()
		custom.MustParse(gormCustomTpl)
		custom.MustExecute(ctx)
		custom.MustSave(customFilename, true)
	}
	return nil
}
