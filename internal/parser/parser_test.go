package parser

import (
	_ "embed"
	"testing"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/stretchr/testify/assert"
)

//go:embed test.sql
var testSql string

func TestParse(t *testing.T) {
	t.Run("ParseError", func(t *testing.T) {
		_, err := Parse("delete from where id = ?")
		assert.NotNil(t, err)
	})

	t.Run("ParseError", func(t *testing.T) {
		_, err := Parse("alter table foo add column name varchar(255);")
		assert.ErrorIs(t, err, errorUnsupportedStmt)
	})

	t.Run("success", func(t *testing.T) {
		dxl, err := Parse(testSql)
		assert.Nil(t, err)

		_, err = spec.From(dxl)
		assert.Nil(t, err)
	})
}
