package parser

import (
	"testing"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/parser/opcode"
	"github.com/stretchr/testify/assert"
)

var testParser *parser.Parser

func TestMain(m *testing.M) {
	testParser = parser.New()
	m.Run()
}

func Test_parseDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stmt, _, err := testParser.Parse(`-- fn: foo
delete from foo where id = ? order by name limit 1`, "", "")
		assert.NoError(t, err)
		for _, v := range stmt {
			deleteStmt, ok := v.(*ast.DeleteStmt)
			if !ok {
				continue
			}
			_, err := parseDelete(deleteStmt)
			assert.NoError(t, err)
		}
	})
	t.Run("parseLineComment", func(t *testing.T) {
		stmt, _, err := testParser.Parse(`-- fn:
delete from foo where id = ?`, "", "")
		assert.NoError(t, err)
		for _, v := range stmt {
			deleteStmt, ok := v.(*ast.DeleteStmt)
			if !ok {
				continue
			}
			_, err := parseDelete(deleteStmt)
			assert.Contains(t, err.Error(), errorMissingFunction.Error())
		}
	})
	t.Run("IsMultiTable", func(t *testing.T) {
		stmt, _, err := testParser.Parse(`-- fn: foo
delete from foo where id = ?`, "", "")
		assert.NoError(t, err)
		for _, v := range stmt {
			deleteStmt, ok := v.(*ast.DeleteStmt)
			if !ok {
				continue
			}
			// mock
			deleteStmt.IsMultiTable = true
			_, err := parseDelete(deleteStmt)
			assert.Contains(t, err.Error(), errorMultipleTable.Error())
		}
	})

	t.Run("parseTableRefsClause", func(t *testing.T) {
		stmt := &ast.DeleteStmt{}
		_, err := parseDelete(stmt)
		assert.Contains(t, err.Error(), errorMissingTable.Error())
	})

	t.Run("whereExpr", func(t *testing.T) {
		stmt := &ast.DeleteStmt{
			Where: &ast.BinaryOperationExpr{
				Op: opcode.Plus,
			},
			TableRefs: &ast.TableRefsClause{
				TableRefs: &ast.Join{
					Left: &ast.TableSource{
						Source: &ast.TableName{
							Name: model.CIStr{
								O: "foo",
								L: "foo",
							},
						},
					},
				},
			},
		}
		_, err := parseDelete(stmt)
		assert.Contains(t, err.Error(), "unsupported opcode")
	})

	t.Run("orderExpr", func(t *testing.T) {
		stmt := &ast.DeleteStmt{
			Order: &ast.OrderByClause{
				Items: []*ast.ByItem{
					{},
				},
			},
			TableRefs: &ast.TableRefsClause{
				TableRefs: &ast.Join{
					Left: &ast.TableSource{
						Source: &ast.TableName{
							Name: model.CIStr{
								O: "foo",
								L: "foo",
							},
						},
					},
				},
			},
		}
		_, err := parseDelete(stmt)
		assert.Contains(t, err.Error(), errorInvalidExprNode.Error())
	})

	t.Run("limitExpr", func(t *testing.T) {
		stmt := &ast.DeleteStmt{
			Limit: &ast.Limit{
				Count: &ast.BetweenExpr{},
			},
			TableRefs: &ast.TableRefsClause{
				TableRefs: &ast.Join{
					Left: &ast.TableSource{
						Source: &ast.TableName{
							Name: model.CIStr{
								O: "foo",
								L: "foo",
							},
						},
					},
				},
			},
		}
		_, err := parseDelete(stmt)
		assert.Contains(t, err.Error(), errorUnsupportedLimitExpr.Error())
	})
}
