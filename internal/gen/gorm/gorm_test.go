package gorm

import (
	_ "embed"
	"log"
	"testing"

	"github.com/anqiansong/sqlgen/internal/parser"
)

//go:embed test.sql
var testSql string

func TestRun(t *testing.T) {
	dxl, err := parser.Parse(testSql)
	if err != nil {
		log.Fatal(err)
	}

	if err := Run(dxl, t.TempDir()); err != nil {
		log.Fatal(err)
	}
}
