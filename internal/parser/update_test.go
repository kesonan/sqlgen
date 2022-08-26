package parser

import (
	"testing"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/parser/opcode"
	"github.com/stretchr/testify/assert"
)

func Test_parseUpdate(t *testing.T) {
	t.Run("missingFunction", func(t *testing.T) {
		stmt, _, err := testParser.Parse(`-- fn:
delete from foo where id = ?`, "", "")
		assert.NoError(t, err)
		for _, v := range stmt {
			updateStmt, ok := v.(*ast.UpdateStmt)
			if !ok {
				continue
			}
			_, err := parseUpdate(updateStmt)
			assert.Contains(t, err.Error(), errorMissingFunction.Error())
		}
	})

	t.Run("parseTableRefsClause", func(t *testing.T) {
		stmt := &ast.UpdateStmt{}
		_, err := parseUpdate(stmt)
		assert.Contains(t, err.Error(), errorMissingTable.Error())
	})

	t.Run("whereExpr", func(t *testing.T) {
		stmt := &ast.UpdateStmt{
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
		_, err := parseUpdate(stmt)
		assert.Contains(t, err.Error(), "unsupported opcode")
	})

	t.Run("orderExpr", func(t *testing.T) {
		stmt := &ast.UpdateStmt{
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
		_, err := parseUpdate(stmt)
		assert.Contains(t, err.Error(), errorInvalidExprNode.Error())
	})

	t.Run("limitExpr", func(t *testing.T) {
		stmt := &ast.UpdateStmt{
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
		_, err := parseUpdate(stmt)
		assert.Contains(t, err.Error(), errorUnsupportedLimitExpr.Error())
	})

}
