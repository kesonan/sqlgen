package parser

import (
	"testing"

	"github.com/pingcap/parser/ast"
	"github.com/stretchr/testify/assert"
)

func Test_parseInsert(t *testing.T) {
	t.Run("missingFunction", func(t *testing.T) {
		stmt, _, err := testParser.Parse(`-- fn:
delete from foo where id = ?`, "", "")
		assert.NoError(t, err)
		for _, v := range stmt {
			insertStmt, ok := v.(*ast.InsertStmt)
			if !ok {
				continue
			}
			_, err := parseInsert(insertStmt)
			assert.Contains(t, err.Error(), errorMissingFunction.Error())
		}
	})

	t.Run("parseTableRefsClause", func(t *testing.T) {
		stmt := &ast.InsertStmt{}
		_, err := parseInsert(stmt)
		assert.Contains(t, err.Error(), errorMissingTable.Error())
	})
}
