package parser

import (
	"errors"
	"testing"

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
}

func Test_parseDSN(t *testing.T) {
	_, _, err := parseDSN("foo")
	assert.NotNil(t, err)

	db, url, err := parseDSN("foo:bar@tcp(127.0.0.1:3306)/foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", db)
	assert.Equal(t, "foo:bar@tcp(127.0.0.1:3306)/information_schema", url)
}
