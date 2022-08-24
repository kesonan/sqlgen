// zeromicro copyright,do not edit.

// MIT License
//
//Copyright (c) 2022 zeromicro
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

// Filepath: go-zero/tools/goctl/model/sql/model/infoschemamodel.go

package infoschema

import (
	"sort"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IInformationSchema = (*InformationSchemaModel)(nil)

type (
	// IInformationSchema defines an interface for schema.
	// Just for mock.
	IInformationSchema interface {
		GetAllTables(database string) ([]string, error)
		FindColumns(db, table string) (*Table, error)
		FindIndex(db, table, column string) ([]*DbIndex, error)
	}

	// InformationSchemaModel defines information schema model
	InformationSchemaModel struct {
		conn sqlx.SqlConn
	}

	// Column defines column in table
	Column struct {
		*DbColumn
		Index *DbIndex
	}

	// DbColumn defines column info of columns
	DbColumn struct {
		Name            string      `db:"COLUMN_NAME"`
		DataType        string      `db:"DATA_TYPE"`
		ColumnType      string      `db:"COLUMN_TYPE"`
		Extra           string      `db:"EXTRA"`
		Comment         string      `db:"COLUMN_COMMENT"`
		ColumnDefault   interface{} `db:"COLUMN_DEFAULT"`
		IsNullAble      string      `db:"IS_NULLABLE"`
		OrdinalPosition int         `db:"ORDINAL_POSITION"`
	}

	// DbIndex defines index of columns in information_schema.statistic
	DbIndex struct {
		IndexName  string `db:"INDEX_NAME"`
		NonUnique  int    `db:"NON_UNIQUE"`
		SeqInIndex int    `db:"SEQ_IN_INDEX"`
	}

	// Table defines table data
	Table struct {
		Db      string
		Table   string
		Columns []*Column
	}
)

// NewInformationSchemaModel creates an instance for InformationSchemaModel
func NewInformationSchemaModel(conn sqlx.SqlConn) IInformationSchema {
	return &InformationSchemaModel{conn: conn}
}

// GetAllTables selects all tables from TABLE_SCHEMA
func (m *InformationSchemaModel) GetAllTables(database string) ([]string, error) {
	query := `select TABLE_NAME from TABLES where TABLE_SCHEMA = ?`
	var tables []string
	err := m.conn.QueryRows(&tables, query, database)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

// FindColumns return columns in specified database and table
func (m *InformationSchemaModel) FindColumns(db, table string) (*Table, error) {
	querySql := `SELECT c.COLUMN_NAME,c.DATA_TYPE,c.COLUMN_TYPE,EXTRA,c.COLUMN_COMMENT,c.COLUMN_DEFAULT,c.IS_NULLABLE,c.ORDINAL_POSITION from COLUMNS c WHERE c.TABLE_SCHEMA = ? and c.TABLE_NAME = ?`
	var reply []*DbColumn
	err := m.conn.QueryRowsPartial(&reply, querySql, db, table)
	if err != nil {
		return nil, err
	}

	var list []*Column
	for _, item := range reply {
		index, err := m.FindIndex(db, table, item.Name)
		if err != nil {
			if err != sqlx.ErrNotFound {
				return nil, err
			}
			continue
		}

		if len(index) > 0 {
			for _, i := range index {
				list = append(list, &Column{
					DbColumn: item,
					Index:    i,
				})
			}
		} else {
			list = append(list, &Column{
				DbColumn: item,
			})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].OrdinalPosition < list[j].OrdinalPosition
	})

	var ret Table
	ret.Db = db
	ret.Table = table
	ret.Columns = list
	return &ret, nil
}

// FindIndex finds index with given db, table and column.
func (m *InformationSchemaModel) FindIndex(db, table, column string) ([]*DbIndex, error) {
	querySql := `SELECT s.INDEX_NAME,s.NON_UNIQUE,s.SEQ_IN_INDEX from  STATISTICS s  WHERE  s.TABLE_SCHEMA = ? and s.TABLE_NAME = ? and s.COLUMN_NAME = ?`
	var reply []*DbIndex
	err := m.conn.QueryRowsPartial(&reply, querySql, db, table, column)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
