package parser

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"strings"
	"text/template"

	sql "github.com/go-sql-driver/mysql"
	"github.com/pingcap/parser/mysql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/anqiansong/sqlgen/internal/infoschema"
	"github.com/anqiansong/sqlgen/internal/patterns"
	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/anqiansong/sqlgen/internal/stringx"
	"github.com/anqiansong/sqlgen/internal/templatex"
)

var errMissingSchema = errors.New("missing schema")

func From(dsn string, pattern ...string) (*spec.DXL, error) {
	schema, url, err := parseDSN(dsn)
	if err != nil {
		return nil, err
	}

	var conn = sqlx.NewMysql(url)
	var model = infoschema.NewInformationSchemaModel(conn)
	tables, err := model.GetAllTables(schema)
	if err != nil {
		return nil, err
	}

	var p = patterns.New(pattern...)
	var matchTables = p.Match(tables...)
	var dxl spec.DXL
	for _, table := range matchTables {
		modelTable, err := model.FindColumns(schema, table)
		if err != nil {
			return nil, err
		}

		ddl, err := convertDDL(modelTable)
		if err != nil {
			return nil, err
		}

		dml, err := convertDML(ddl.Table)
		if err != nil {
			return nil, err
		}

		dxl.DDL = append(dxl.DDL, ddl)
		dxl.DML = append(dxl.DML, dml...)
	}

	return &dxl, nil
}

func parseDSN(dsn string) (db, url string, err error) {
	cfg, err := sql.ParseDSN(dsn)
	if err != nil {
		return "", "", err
	}

	if cfg.DBName == "" {
		return "", "", errMissingSchema
	}

	url = fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.User, cfg.Passwd, cfg.Addr, "information_schema")
	db = cfg.DBName
	return
}

//go:embed init.tpl.sql
var initSql string

func convertDML(in *spec.Table) ([]spec.DML, error) {
	t, err := template.New("sql").Parse(initSql)
	if err != nil {
		return nil, err
	}

	var sqlBuffer bytes.Buffer
	if err = t.Execute(&sqlBuffer, map[string]interface{}{
		"insert_columns": strings.Join(in.ColumnList(), ", "),
		"insert_table":   in.Name,
		"insert_values":  stringx.RepeatJoin("?", ", ", len(in.ColumnList())),
		"unique_indexes": getUniques(in),
	}); err != nil {
		return nil, err
	}

	dxl, err := Parse(sqlBuffer.String())
	if err != nil {
		return nil, err
	}

	return dxl.DML, nil
}

// Unique is a unique index info.
type Unique struct {
	SelectColumns  string
	Table          string
	UpdateSet      string
	WhereClause    string
	UniqueNameJoin string
}

func getUniques(in *spec.Table) []Unique {
	var list []Unique
	var columns = strings.Join(in.ColumnList(), ", ")
	var updateSet = strings.Join(in.ColumnList(), " = ?,") + " = ?"
	var m = map[Unique]struct{}{}
	for _, c := range in.Constraint.PrimaryKey {
		var item = Unique{
			SelectColumns:  columns,
			Table:          in.Name,
			UpdateSet:      updateSet,
			WhereClause:    strings.Join(c, " = ? AND") + " = ?",
			UniqueNameJoin: templatex.UpperCamel(strings.Join(c, "")),
		}
		if _, ok := m[item]; ok {
			continue
		}
		m[item] = struct{}{}
		list = append(list, item)
	}
	for _, c := range in.Constraint.UniqueKey {
		var item = Unique{
			SelectColumns:  columns,
			Table:          in.Name,
			UpdateSet:      updateSet,
			WhereClause:    strings.Join(c, " = ? AND") + " = ?",
			UniqueNameJoin: templatex.UpperCamel(strings.Join(c, "")),
		}
		if _, ok := m[item]; ok {
			continue
		}
		m[item] = struct{}{}
		list = append(list, item)
	}

	return list
}

func convertDDL(in *infoschema.Table) (*spec.DDL, error) {
	var ddl spec.DDL
	var constraint = spec.NewConstraint()
	getConstraint(in.Columns, constraint)
	var table spec.Table
	table.Name = in.Table
	table.Schema = in.Db
	if !constraint.IsEmpty() {
		table.Constraint = *constraint
	}

	for _, c := range in.Columns {
		var extra = c.Extra
		var autoIncrement = strings.Contains(extra, "auto_increment")
		var unsigned = strings.Contains(c.DataType, "unsigned")
		tp, err := dbTypeMapper(c.DataType)
		if err != nil {
			return nil, err
		}

		table.Columns = append(table.Columns, spec.Column{
			ColumnOption: spec.ColumnOption{
				AutoIncrement:   autoIncrement,
				Comment:         stringx.TrimNewLine(c.Comment),
				HasDefaultValue: c.ColumnDefault != nil,
				NotNull:         !strings.EqualFold(c.IsNullAble, "yes"),
				Unsigned:        unsigned,
			},
			Name: c.Name,
			TP:   tp,
		})
	}

	ddl.Table = &table
	return &ddl, nil
}

func getConstraint(columns []*infoschema.Column, constraint *spec.Constraint) {
	for _, c := range columns {
		index := c.Index
		if index == nil {
			continue
		}
		indexName := index.IndexName
		if strings.EqualFold(indexName, "primary") {
			constraint.AppendPrimaryKey(indexName, c.Name)
		}
		if index.NonUnique == 0 {
			constraint.AppendUniqueKey(indexName, c.Name)
		} else {
			constraint.AppendIndex(indexName, c.Name)
		}
	}
}

var str2Type = map[string]byte{
	"bit":         mysql.TypeBit,
	"text":        mysql.TypeBlob,
	"date":        mysql.TypeDate,
	"datetime":    mysql.TypeDatetime,
	"unspecified": mysql.TypeUnspecified,
	"decimal":     mysql.TypeNewDecimal,
	"double":      mysql.TypeDouble,
	"enum":        mysql.TypeEnum,
	"float":       mysql.TypeFloat,
	"geometry":    mysql.TypeGeometry,
	"mediumint":   mysql.TypeInt24,
	"json":        mysql.TypeJSON,
	"int":         mysql.TypeLong,
	"bigint":      mysql.TypeLonglong,
	"longtext":    mysql.TypeLongBlob,
	"mediumtext":  mysql.TypeMediumBlob,
	"null":        mysql.TypeNull,
	"set":         mysql.TypeSet,
	"smallint":    mysql.TypeShort,
	"char":        mysql.TypeString,
	"time":        mysql.TypeDuration,
	"timestamp":   mysql.TypeTimestamp,
	"tinyint":     mysql.TypeTiny,
	"tinytext":    mysql.TypeTinyBlob,
	"varchar":     mysql.TypeVarchar,
	"var_string":  mysql.TypeVarString,
	"year":        mysql.TypeYear,
}

func dbTypeMapper(tp string) (byte, error) {
	var l = strings.ToLower(tp)
	ret, ok := str2Type[l]
	if !ok {
		return 0, fmt.Errorf("unsupported type:%s", tp)
	}
	return ret, nil
}
