package parser

import (
	"testing"

	"github.com/pingcap/parser/ast"
	"github.com/stretchr/testify/assert"
)

func Test_parseDML(t *testing.T) {
	t.Run("InsertStmt", func(t *testing.T) {
		stmt, _, err := testParser.Parse(`
-- fn:
insert into foo (name) values(?);
-- fn:
select * from foo;
-- fn:
delete from foo where id = ?;
-- fn:
update foo set name = ? where id = ?;
-- fn:
alter table foo add column bar varchar(255);
`, "", "")
		assert.NoError(t, err)
		for _, v := range stmt {
			_, err := parseDML(v)
			assert.NotNil(t, err)
		}
	})
}

func Test_parseTableRefsClause(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		_, err := parseTableRefsClause(nil)
		assert.ErrorIs(t, err, errorMissingTable)
	})

	t.Run("joinNil", func(t *testing.T) {
		_, err := parseTableRefsClause(&ast.TableRefsClause{})
		assert.ErrorIs(t, err, errorMissingTable)
	})

	t.Run("joinLeftNil", func(t *testing.T) {
		_, err := parseTableRefsClause(&ast.TableRefsClause{
			TableRefs: &ast.Join{
				Left:  &ast.TableName{},
				Right: &ast.TableName{},
			},
		})
		assert.ErrorIs(t, err, errorMultipleTable)
	})

	t.Run("parseResultSetNode", func(t *testing.T) {
		_, err := parseTableRefsClause(&ast.TableRefsClause{
			TableRefs: &ast.Join{
				Left: &ast.SelectStmt{},
			},
		})
		assert.ErrorIs(t, err, errorUnsupportedNestedQuery)
	})
}

func Test_parseTransaction(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		_, err := parseTransaction(nil)
		assert.ErrorIs(t, err, errorMissingTransaction)
	})
}
