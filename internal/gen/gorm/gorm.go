package gorm

import (
	_ "embed"
	"os"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/anqiansong/sqlgen/internal/templatex"
)

//go:embed gorm.tpl
var gormTpl string

func Run(dxl *spec.DXL) error {
	var t = templatex.New()
	for _, ddl := range dxl.DDL {
		t.AppendFuncMap(map[string]any{
			"IsPrimary": ddl.Table.IsPrimary,
		})
		t.MustParse(gormTpl)
		insertStmts, updateStmts, selectStmts, deleteStmts := dxl.Stmt(ddl.Table.Name)
		t.MustExecute(map[string]interface{}{
			"Table":      ddl.Table,
			"InsertStmt": insertStmts,
			"SelectStmt": updateStmts,
			"DeleteStmt": selectStmts,
			"UpdateStmt": deleteStmts,
		})
		t.Write(os.Stdout, true)
	}

	return nil
}
