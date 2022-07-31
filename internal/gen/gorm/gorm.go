package gorm

import (
	_ "embed"

	"github.com/anqiansong/sqlgen/internal/spec"
)

//go:embed gorm.tpl
var gormTpl string

func Run(dxl *spec.DXL) error {
	return nil
}
