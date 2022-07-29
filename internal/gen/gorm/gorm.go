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
	t, err := template.New("gorm").Funcs(
		template.FuncMap{
			"UpperCamel": templatex.UpperCamel,
			"LowerCamel": templatex.LowerCamel,
		}).Parse(gormTpl)
	if err != nil {
		return err
	}

	for _, ddl := range dxl.DDL {
		if err = t.Execute(os.Stdout, ddl.Table); err != nil {
			return err
		}
	}

	return nil
}
