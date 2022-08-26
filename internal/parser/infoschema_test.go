package parser

import (
	"errors"
	"testing"
	"text/template"

	gomonkey "github.com/agiledragon/gomonkey/v2"
	"github.com/anqiansong/sqlgen/internal/infoschema"
	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/stretchr/testify/assert"
)

var dummyError = errors.New("dummy")

func TestFrom(t *testing.T) {
	t.Run("invalidDSN", func(t *testing.T) {
		_, err := From("foo")
		assert.NotNil(t, err)
	})
	t.Run("GetAllTables", func(t *testing.T) {
		model := infoschema.NewInformationSchemaModel(nil)
		patch := gomonkey.ApplyFunc(parseDSN, func(database string) (string, string, error) {
			return "", "", nil
		})
		patch.ApplyMethodFunc(model, "GetAllTables", func(database string) ([]string, error) {
			return nil, dummyError
		})
		t.Cleanup(func() {
			patch.Reset()
		})
		_, err := From("foo")
		assert.ErrorIs(t, err, dummyError)
	})
	t.Run("FindColumns", func(t *testing.T) {
		model := infoschema.NewInformationSchemaModel(nil)
		patch := gomonkey.ApplyFunc(parseDSN, func(database string) (string, string, error) {
			return "", "", nil
		})
		patch.ApplyMethodFunc(model, "GetAllTables", func(database string) ([]string, error) {
			return []string{"foo", "bar"}, nil
		})
		patch.ApplyMethodFunc(model, "FindColumns", func(db, table string) (*infoschema.Table, error) {
			return nil, dummyError
		})
		t.Cleanup(func() {
			patch.Reset()
		})
		_, err := From("foo")
		assert.ErrorIs(t, err, dummyError)
	})

	t.Run("convertDDL", func(t *testing.T) {
		model := infoschema.NewInformationSchemaModel(nil)
		patch := gomonkey.ApplyFunc(parseDSN, func(database string) (string, string, error) {
			return "", "", nil
		})
		patch.ApplyMethodFunc(model, "GetAllTables", func(database string) ([]string, error) {
			return []string{"foo", "bar"}, nil
		})
		patch.ApplyMethodFunc(model, "FindColumns", func(db, table string) (*infoschema.Table, error) {
			return &infoschema.Table{}, nil
		})
		patch.ApplyFunc(convertDDL, func(in *infoschema.Table) (*spec.DDL, error) {
			return nil, dummyError
		})
		t.Cleanup(func() {
			patch.Reset()
		})
		_, err := From("foo")
		assert.ErrorIs(t, err, dummyError)
	})

	t.Run("convertDDL", func(t *testing.T) {
		model := infoschema.NewInformationSchemaModel(nil)
		patch := gomonkey.ApplyFunc(parseDSN, func(database string) (string, string, error) {
			return "", "", nil
		})
		patch.ApplyMethodFunc(model, "GetAllTables", func(database string) ([]string, error) {
			return []string{"foo", "bar"}, nil
		})
		patch.ApplyMethodFunc(model, "FindColumns", func(db, table string) (*infoschema.Table, error) {
			return &infoschema.Table{}, nil
		})
		patch.ApplyFunc(convertDDL, func(in *infoschema.Table) (*spec.DDL, error) {
			return &spec.DDL{}, nil
		})
		patch.ApplyFunc(convertDML, func(in *spec.Table) ([]spec.DML, error) {
			return nil, dummyError
		})
		t.Cleanup(func() {
			patch.Reset()
		})
		_, err := From("foo")
		assert.ErrorIs(t, err, dummyError)
	})

	t.Run("success", func(t *testing.T) {
		model := infoschema.NewInformationSchemaModel(nil)
		patch := gomonkey.ApplyFunc(parseDSN, func(database string) (string, string, error) {
			return "", "", nil
		})
		patch.ApplyMethodFunc(model, "GetAllTables", func(database string) ([]string, error) {
			return []string{"foo", "bar"}, nil
		})
		patch.ApplyMethodFunc(model, "FindColumns", func(db, table string) (*infoschema.Table, error) {
			return &infoschema.Table{}, nil
		})
		patch.ApplyFunc(convertDDL, func(in *infoschema.Table) (*spec.DDL, error) {
			return &spec.DDL{}, nil
		})
		patch.ApplyFunc(convertDML, func(in *spec.Table) ([]spec.DML, error) {
			return []spec.DML{}, nil
		})
		t.Cleanup(func() {
			patch.Reset()
		})
		_, err := From("foo")
		assert.Nil(t, err)
	})
}

func Test_parseDSN(t *testing.T) {
	_, _, err := parseDSN("foo")
	assert.NotNil(t, err)

	db, url, err := parseDSN("foo:bar@tcp(127.0.0.1:3306)/foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", db)
	assert.Equal(t, "foo:bar@tcp(127.0.0.1:3306)/information_schema", url)
}

func Test_convertDML(t *testing.T) {
	t.Run("parseError", func(t *testing.T) {
		tpl := template.New("foo")
		patch := gomonkey.ApplyMethodFunc(tpl, "Parse", func(text string) (*template.Template, error) {
			return nil, dummyError
		})
		t.Cleanup(func() {
			patch.Reset()
		})
		_, err := convertDML(nil)
		assert.ErrorIs(t, err, dummyError)
	})

	t.Run("Parse", func(t *testing.T) {
		patch := gomonkey.ApplyFunc(Parse, func(sql string) (*spec.DXL, error) {
			return nil, dummyError
		})
		t.Cleanup(func() {
			patch.Reset()
		})
		_, err := convertDML(&spec.Table{
			Columns: spec.Columns{
				{Name: "foo"},
				{Name: "bar"},
			},
			Name: "foo",
		})
		assert.ErrorIs(t, err, dummyError)
	})

	t.Run("success", func(t *testing.T) {
		_, err := convertDML(&spec.Table{
			Columns: spec.Columns{
				{Name: "foo"},
				{Name: "bar"},
			},
			Name: "foo",
		})
		assert.Nil(t, err)
	})
}

func Test_getUniques(t *testing.T) {
	unique := getUniques(&spec.Table{
		Columns: spec.Columns{
			{Name: "foo"},
			{Name: "foo"},
			{Name: "baz"},
		},
		Constraint: spec.Constraint{
			PrimaryKey: map[string][]string{
				"foo": {"bar"},
				"baz": {"bar"},
			},
			UniqueKey: map[string][]string{
				"bar": {"baz"},
				"baz": {"baz"},
			},
		},
		Schema: "",
		Name:   "foo",
	})

	assert.Equal(t, 2, len(unique))
}

func Test_convertDDL(t *testing.T) {
	t.Run("constraint.IsEmpty", func(t *testing.T) {
		_, err := convertDDL(&infoschema.Table{})
		assert.Nil(t, err)
	})

	t.Run("!constraint.IsEmpty", func(t *testing.T) {
		_, err := convertDDL(&infoschema.Table{
			Db:    "foo",
			Table: "foo",
			Columns: []*infoschema.Column{
				{
					DbColumn: &infoschema.DbColumn{
						Name:            "foo",
						DataType:        "bit",
						ColumnType:      "bit",
						Extra:           "auto_increment",
						Comment:         "foo",
						IsNullAble:      "yes",
						OrdinalPosition: 0,
					},
					Index: &infoschema.DbIndex{
						IndexName: "foo",
					},
				},
			},
		})
		assert.Nil(t, err)
	})

	t.Run("dbTypeMapper", func(t *testing.T) {
		_, err := convertDDL(&infoschema.Table{
			Db:    "foo",
			Table: "bar",
			Columns: []*infoschema.Column{
				{
					DbColumn: &infoschema.DbColumn{
						Name:            "foo",
						DataType:        "foo",
						ColumnType:      "bit",
						Extra:           "auto_increment",
						Comment:         "foo",
						IsNullAble:      "yes",
						OrdinalPosition: 0,
					},
					Index: &infoschema.DbIndex{
						IndexName: "foo",
					},
				},
			},
		})
		assert.Contains(t, err.Error(), "unsupported type")
	})

	t.Run("success", func(t *testing.T) {
		_, err := convertDDL(&infoschema.Table{
			Db:    "foo",
			Table: "bar",
			Columns: []*infoschema.Column{
				{
					DbColumn: &infoschema.DbColumn{
						Name:            "foo",
						DataType:        "bit",
						ColumnType:      "bit",
						Extra:           "auto_increment",
						Comment:         "foo",
						IsNullAble:      "yes",
						OrdinalPosition: 0,
					},
				},
				{
					DbColumn: &infoschema.DbColumn{
						Name:            "foo",
						DataType:        "bit",
						ColumnType:      "bit",
						Extra:           "auto_increment",
						Comment:         "foo",
						IsNullAble:      "yes",
						ColumnDefault:   "",
						OrdinalPosition: 1,
					},
				},
			},
		})
		assert.Nil(t, err)
	})
}

func Test_getConstraint(t *testing.T) {
	var constraint = spec.NewConstraint()
	getConstraint([]*infoschema.Column{
		{
			DbColumn: &infoschema.DbColumn{
				Name:            "foo",
				DataType:        "foo",
				ColumnType:      "bit",
				Extra:           "auto_increment",
				Comment:         "foo",
				IsNullAble:      "yes",
				OrdinalPosition: 0,
			},
			Index: nil,
		},
		{
			DbColumn: &infoschema.DbColumn{
				Name:            "foo",
				DataType:        "foo",
				ColumnType:      "bit",
				Extra:           "auto_increment",
				Comment:         "foo",
				IsNullAble:      "yes",
				OrdinalPosition: 0,
			},
			Index: &infoschema.DbIndex{
				IndexName: "foo",
			},
		},
		{
			DbColumn: &infoschema.DbColumn{
				Name:            "foo",
				DataType:        "foo",
				ColumnType:      "bit",
				Extra:           "auto_increment",
				Comment:         "foo",
				IsNullAble:      "yes",
				OrdinalPosition: 0,
			},
			Index: &infoschema.DbIndex{
				IndexName:  "primary",
				SeqInIndex: 1,
			},
		},
		{
			DbColumn: &infoschema.DbColumn{
				Name:            "foo",
				DataType:        "foo",
				ColumnType:      "bit",
				Extra:           "auto_increment",
				Comment:         "foo",
				IsNullAble:      "yes",
				OrdinalPosition: 0,
			},
			Index: &infoschema.DbIndex{
				NonUnique:  1,
				SeqInIndex: 2,
			},
		},
	}, constraint)
}
