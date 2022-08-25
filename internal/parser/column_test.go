package parser

import (
	"testing"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/test_driver"
	"github.com/pingcap/parser/types"
	"github.com/stretchr/testify/assert"
)

func Test_parseColumnDef(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		column, constraint := parseColumnDef(nil)
		assert.Nil(t, column)
		assert.Nil(t, constraint)

		column, constraint = parseColumnDef(&ast.ColumnDef{})
		assert.Nil(t, column)
		assert.Nil(t, constraint)
	})

	t.Run("not nil", func(t *testing.T) {
		datum := test_driver.Datum{}
		datum.SetString("foo")
		name := &ast.ColumnName{
			Name: model.CIStr{
				O: "foo",
				L: "foo",
			},
		}
		tp := &types.FieldType{
			Tp:   mysql.TypeBit,
			Flag: mysql.UnsignedFlag,
		}
		column, constraint := parseColumnDef(&ast.ColumnDef{
			Name: name,
			Tp:   tp,
			Options: []*ast.ColumnOption{
				{Tp: ast.ColumnOptionNotNull},
				{Tp: ast.ColumnOptionAutoIncrement},
				{Tp: ast.ColumnOptionDefaultValue},
				{
					Tp:   ast.ColumnOptionComment,
					Expr: &test_driver.ValueExpr{Datum: datum},
				},
				{Tp: ast.ColumnOptionUniqKey},
				{Tp: ast.ColumnOptionPrimaryKey},
				{Tp: ast.ColumnOptionFulltext},
			},
		})
		assert.Equal(t, spec.Column{
			ColumnOption: spec.ColumnOption{
				AutoIncrement:   true,
				Comment:         "foo",
				HasDefaultValue: true,
				NotNull:         true,
				Unsigned:        true,
			},
			Name:          "foo",
			TP:            mysql.TypeBit,
			AggregateCall: false,
		}, *column)

		assertMapEqual(t, map[string][]string{
			"foo": {"foo"},
		}, constraint.UniqueKey)
		assertMapEqual(t, map[string][]string{
			"foo": {"foo"},
		}, constraint.PrimaryKey)
	})
}

func assertMapEqual(t *testing.T, m1, m2 map[string][]string) {
	for k, v1 := range m1 {
		v2, ok := m2[k]
		assert.True(t, ok)
		assert.Equal(t, v1, v2)
	}
	return
}

func Test_parseConstraint(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		c := parseConstraint(nil)
		assert.Nil(t, c)

		c = parseConstraint(&ast.Constraint{Keys: nil})
		assert.Nil(t, c)
	})

	t.Run("ConstraintPrimaryKey", func(t *testing.T) {
		c := parseConstraint(&ast.Constraint{
			Name: "foo",
			Tp:   ast.ConstraintPrimaryKey,
			Keys: []*ast.IndexPartSpecification{
				{
					Column: &ast.ColumnName{
						Name: model.CIStr{
							O: "foo",
							L: "foo",
						},
					},
				},
			}})
		assertMapEqual(t, map[string][]string{
			"foo": {"foo"},
		}, c.PrimaryKey)
	})

	t.Run("ConstraintKey", func(t *testing.T) {
		c := parseConstraint(&ast.Constraint{
			Name: "foo",
			Tp:   ast.ConstraintKey,
			Keys: []*ast.IndexPartSpecification{
				{
					Column: &ast.ColumnName{
						Name: model.CIStr{
							O: "foo",
							L: "foo",
						},
					},
				},
			}})
		assertMapEqual(t, map[string][]string{
			"foo": {"foo"},
		}, c.Index)
	})

	t.Run("ConstraintIndex", func(t *testing.T) {
		c := parseConstraint(&ast.Constraint{
			Name: "foo",
			Tp:   ast.ConstraintIndex,
			Keys: []*ast.IndexPartSpecification{
				{
					Column: &ast.ColumnName{
						Name: model.CIStr{
							O: "foo",
							L: "foo",
						},
					},
				},
			}})
		assertMapEqual(t, map[string][]string{
			"foo": {"foo"},
		}, c.Index)
	})

	t.Run("ConstraintUniq", func(t *testing.T) {
		c := parseConstraint(&ast.Constraint{
			Name: "foo",
			Tp:   ast.ConstraintUniq,
			Keys: []*ast.IndexPartSpecification{
				{
					Column: &ast.ColumnName{
						Name: model.CIStr{
							O: "foo",
							L: "foo",
						},
					},
				},
			}})
		assertMapEqual(t, map[string][]string{
			"foo": {"foo"},
		}, c.UniqueKey)
	})
	t.Run("ConstraintUniqKey", func(t *testing.T) {
		c := parseConstraint(&ast.Constraint{
			Name: "foo",
			Tp:   ast.ConstraintUniqKey,
			Keys: []*ast.IndexPartSpecification{
				{
					Column: &ast.ColumnName{
						Name: model.CIStr{
							O: "foo",
							L: "foo",
						},
					},
				},
			}})
		assertMapEqual(t, map[string][]string{
			"foo": {"foo"},
		}, c.UniqueKey)
	})
	t.Run("ConstraintUniqIndex", func(t *testing.T) {
		c := parseConstraint(&ast.Constraint{
			Name: "foo",
			Tp:   ast.ConstraintUniqIndex,
			Keys: []*ast.IndexPartSpecification{
				{
					Column: &ast.ColumnName{
						Name: model.CIStr{
							O: "foo",
							L: "foo",
						},
					},
				},
			}})
		assertMapEqual(t, map[string][]string{
			"foo": {"foo"},
		}, c.UniqueKey)
	})

	t.Run("ConstraintUniqIndex", func(t *testing.T) {
		c := parseConstraint(&ast.Constraint{
			Name: "foo",
			Tp:   ast.ConstraintFulltext,
			Keys: []*ast.IndexPartSpecification{
				{
					Column: &ast.ColumnName{
						Name: model.CIStr{
							O: "foo",
							L: "foo",
						},
					},
				},
			}})
		assert.Equal(t, spec.Constraint{
			Index:      map[string][]string{},
			PrimaryKey: map[string][]string{},
			UniqueKey:  map[string][]string{},
		}, *c)
	})
}

func Test_parseColumnFromKeys(t *testing.T) {
	var testData = []struct {
		input  string
		expect []string
	}{
		{
			input:  "foo",
			expect: []string{"foo"},
		},
		{
			expect: []string(nil),
		},
	}
	getColumn := func(name string) *ast.ColumnName {
		if len(name) == 0 {
			return nil
		}
		var column ast.ColumnName
		column.Name = model.CIStr{
			O: name,
			L: name,
		}
		return &column
	}
	for _, v := range testData {
		actual := parseColumnFromKeys([]*ast.IndexPartSpecification{
			{
				Column: getColumn(v.input),
			},
		})
		assert.Equal(t, v.expect, actual)
	}
}
