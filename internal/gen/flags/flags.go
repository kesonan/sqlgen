package flags

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/anqiansong/sqlgen/internal/gen/gorm"
	"github.com/anqiansong/sqlgen/internal/gen/sql"
	"github.com/anqiansong/sqlgen/internal/gen/sqlx"
	"github.com/anqiansong/sqlgen/internal/gen/xorm"
	"github.com/anqiansong/sqlgen/internal/log"
	"github.com/anqiansong/sqlgen/internal/parser"
	"github.com/anqiansong/sqlgen/internal/patterns"
	"github.com/anqiansong/sqlgen/internal/spec"
)

const sqlExt = ".sql"

type Mode int

const (
	SQL Mode = iota
	GORM
	XORM
	SQLX
)

func Run(dsn string, filename, table []string, mode Mode) {
	var err error
	if len(filename) > 0 {
		err = runFromSQL(filename, mode)
	} else if len(dsn) > 0 {
		err = runFromDSN(dsn, table, mode)
	} else {
		err = fmt.Errorf("missing dsn or filename")
	}
	log.Must(err)
}

func runFromSQL(filenamePatterns []string, mode Mode) error {
	var list []string
	for _, filename := range filenamePatterns {
		var dir = filepath.Dir(filename)
		var base = filepath.Base(filename)
		fileInfo, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}
		var filenames []string
		for _, item := range fileInfo {
			if item.IsDir() {
				continue
			}
			ext := filepath.Ext(item.Name())
			if ext != sqlExt {
				continue
			}

			var f = filepath.Join(dir, item.Name())
			filenames = append(filenames, f)
		}
		var p = patterns.New(base)
		var matchSQLFile = p.Match(filenames...)

		list = append(list, matchSQLFile...)
	}

	if len(list) == 0 {
		return fmt.Errorf("no sql file found")
	}

	var ret spec.DXL
	for _, file := range list {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		dxl, err := parser.Parse(string(data))
		if err != nil {
			return err
		}

		ret.DDL = append(ret.DDL, dxl.DDL...)
		ret.DML = append(ret.DML, dxl.DML...)
	}

	return run(&ret, mode)
}

func runFromDSN(dsn string, patterns []string, mode Mode) error {
	dxl, err := parser.From(dsn, patterns...)
	if err != nil {
		return err
	}

	return run(dxl, mode)
}

func run(dxl *spec.DXL, mode Mode) error {
	switch mode {
	case SQL:
		return sql.Run(dxl)
	case GORM:
		return gorm.Run(dxl)
	case XORM:
		return xorm.Run(dxl)
	case SQLX:
		return sqlx.Run(dxl)
	}
	return nil
}
