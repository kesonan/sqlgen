package gorm

import (
	_ "embed"
	"testing"

	"github.com/anqiansong/sqlgen/internal/gen/testdata"
	"github.com/anqiansong/sqlgen/internal/parser"
	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	dxl, err := parser.Parse(testdata.TestSql)
	assert.NoError(t, err)
	ctx, err := spec.From(dxl)
	assert.NoError(t, err)
	err = Run(ctx, t.TempDir())
	assert.NoError(t, err)
}
