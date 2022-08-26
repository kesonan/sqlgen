package parser

import (
	"testing"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/parser/opcode"
	"github.com/stretchr/testify/assert"
)

func Test_parseSelect(t *testing.T) {
	t.Run("missingFunction", func(t *testing.T) {
		stmt, _, err := testParser.Parse(`-- fn:
delete from foo where id = ?`, "", "")
		assert.NoError(t, err)
		for _, v := range stmt {
			selectStmt, ok := v.(*ast.SelectStmt)
			if !ok {
				continue
			}
			_, err := parseSelect(selectStmt)
			assert.Contains(t, err.Error(), errorMissingFunction.Error())
		}
	})

	t.Run("parseTableRefsClause", func(t *testing.T) {
		stmt := &ast.SelectStmt{}
		_, err := parseSelect(stmt)
		assert.Contains(t, err.Error(), errorMissingTable.Error())
	})

	t.Run("whereExpr", func(t *testing.T) {
		stmt := &ast.SelectStmt{
			Where: &ast.BinaryOperationExpr{
				Op: opcode.Plus,
			},
			From: &ast.TableRefsClause{
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
		_, err := parseSelect(stmt)
		assert.Contains(t, err.Error(), "unsupported opcode")
	})

	t.Run("groupBy", func(t *testing.T) {
		stmt := &ast.SelectStmt{
			GroupBy: &ast.GroupByClause{
				Items: []*ast.ByItem{
					{},
				},
			},
			From: &ast.TableRefsClause{
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
		_, err := parseSelect(stmt)
		assert.Contains(t, err.Error(), errorInvalidExprNode.Error())
	})

	t.Run("having", func(t *testing.T) {
		stmt := &ast.SelectStmt{
			Having: &ast.HavingClause{
				Expr: &ast.BinaryOperationExpr{
					Op: opcode.Plus,
				},
			},
			From: &ast.TableRefsClause{
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
		_, err := parseSelect(stmt)
		assert.Contains(t, err.Error(), "unsupported opcode")
	})

	t.Run("orderExpr", func(t *testing.T) {
		stmt := &ast.SelectStmt{
			OrderBy: &ast.OrderByClause{
				Items: []*ast.ByItem{
					{},
				},
			},
			From: &ast.TableRefsClause{
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
		_, err := parseSelect(stmt)
		assert.Contains(t, err.Error(), errorInvalidExprNode.Error())
	})

	t.Run("limitExpr", func(t *testing.T) {
		stmt := &ast.SelectStmt{
			Limit: &ast.Limit{
				Count: &ast.BetweenExpr{},
			},
			From: &ast.TableRefsClause{
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
		_, err := parseSelect(stmt)
		assert.Contains(t, err.Error(), errorUnsupportedLimitExpr.Error())
	})

	t.Run("fields", func(t *testing.T) {
		stmt := &ast.SelectStmt{
			Fields: &ast.FieldList{
				Fields: []*ast.SelectField{
					{
						Offset: 0,
						WildCard: &ast.WildCardField{
							Table: model.CIStr{
								O: "bar",
								L: "bar",
							},
						},
					},
				},
			},
			From: &ast.TableRefsClause{
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
		_, err := parseSelect(stmt)
		assert.Contains(t, err.Error(), "wildcard table")
	})
}

func Test_convertOP(t *testing.T) {
	_, err := convertOP(opcode.In)
	assert.NoError(t, err)

	_, err = convertOP(opcode.Regexp)
	assert.NotNil(t, err)
}
