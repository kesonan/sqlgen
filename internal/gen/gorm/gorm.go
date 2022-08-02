package gorm

import (
	_ "embed"
	"os"
	"text/template"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/anqiansong/sqlgen/internal/templatex"
)

//go:embed gorm.tpl
var gormTpl string

func Run(dxl *spec.DXL) error {
	list, err := spec.From(dxl)
	if err != nil {
		return err
	}

	for _, ctx := range list {
		t := templatex.New()
		t.AppendFuncMap(template.FuncMap{
			"IsPrimary": func(name string) bool {
				return ctx.Table.IsPrimary(name)
			},
		})
		t.MustParse(gormTpl)
		t.MustExecute(ctx)
		t.Write(os.Stdout, true)
	}
	return nil
}
