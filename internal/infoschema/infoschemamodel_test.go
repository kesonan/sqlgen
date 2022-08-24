package infoschema

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var dummyError = errors.New("dummy")

func TestNewInformationSchemaModel(t *testing.T) {
	conn := sqlx.NewMysql("foo")
	instance := NewInformationSchemaModel(conn)
	assert.NotNil(t, instance)
}

func TestInformationSchemaModel_GetAllTables(t *testing.T) {
	logx.Disable()
	var query = `select TABLE_NAME from TABLES where TABLE_SCHEMA = ?`
	var database = "foo"
	var mockTableNames = []string{"foo"}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	mock.ExpectQuery(query).WithArgs(database).WillReturnRows(sqlmock.NewRows(mockTableNames))

	conn := sqlx.NewSqlConnFromDB(db)
	model := NewInformationSchemaModel(conn)
	_, err = model.GetAllTables(database)
	assert.NoError(t, err)

	mock.ExpectQuery(query).WithArgs(database).WillReturnError(dummyError)
	_, err = model.GetAllTables(database)
	assert.ErrorIs(t, err, dummyError)
}

func TestInformationSchemaModel_FindColumns(t *testing.T) {
	logx.Disable()
	var indexQuery = `SELECT s.INDEX_NAME,s.NON_UNIQUE,s.SEQ_IN_INDEX from  STATISTICS s  WHERE  s.TABLE_SCHEMA = ? and s.TABLE_NAME = ? and s.COLUMN_NAME = ?`
	var query = `SELECT c.COLUMN_NAME,c.DATA_TYPE,c.COLUMN_TYPE,EXTRA,c.COLUMN_COMMENT,c.COLUMN_DEFAULT,c.IS_NULLABLE,c.ORDINAL_POSITION from COLUMNS c WHERE c.TABLE_SCHEMA = ? and c.TABLE_NAME = ?`
	var database = "foo"
	var table = "bar"
	var column = "baz"
	var mockTableNames = []string{"foo"}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery(query).WithArgs(database, table).WillReturnRows(sqlmock.NewRows(mockTableNames))
	mock.ExpectQuery(indexQuery).WithArgs(database, table, column).WillReturnRows(sqlmock.NewRows(mockTableNames))

	conn := sqlx.NewSqlConnFromDB(db)
	model := NewInformationSchemaModel(conn)
	_, err = model.FindColumns(database, table)
	assert.NoError(t, err)

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	conn = sqlx.NewSqlConnFromDB(db)
	model = NewInformationSchemaModel(conn)
	mock.ExpectQuery(query).WithArgs(database, table).WillReturnError(dummyError)
	_, err = model.FindColumns(database, table)
	assert.ErrorIs(t, err, dummyError)
}

func TestInformationSchemaModel_FindIndex(t *testing.T) {
	logx.Disable()
	var indexQuery = `SELECT s.INDEX_NAME,s.NON_UNIQUE,s.SEQ_IN_INDEX from  STATISTICS s  WHERE  s.TABLE_SCHEMA = ? and s.TABLE_NAME = ? and s.COLUMN_NAME = ?`
	var database = "foo"
	var table = "bar"
	var column = "baz"
	var mockTableNames = []string{"foo"}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery(indexQuery).WithArgs(database, table, column).WillReturnRows(sqlmock.NewRows(mockTableNames))

	conn := sqlx.NewSqlConnFromDB(db)
	model := NewInformationSchemaModel(conn)
	_, err = model.FindIndex(database, table, column)
	assert.NoError(t, err)

	mock.ExpectQuery(indexQuery).WithArgs(database, table, column).WillReturnError(dummyError)
	_, err = model.FindIndex(database, table, column)
	assert.ErrorIs(t, err, dummyError)
}
