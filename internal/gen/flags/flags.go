package flags

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/anqiansong/sqlgen/internal/gen/bun"
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
	BUN
)

type RunArg struct {
	DSN             string
	Filename, Table []string
	Mode            Mode
	Output          string
}

func Run(arg RunArg) {
	var err error
	if len(arg.DSN) > 0 {
		err = runFromDSN(arg)
	} else if len(arg.Filename) > 0 {
		err = runFromSQL(arg)
	} else {
		err = fmt.Errorf("missing dsn or filename")
	}
	log.Must(err)
}

func runFromSQL(arg RunArg) error {
	var list []string
	for _, filename := range arg.Filename {
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

	return run(&ret, arg.Mode, arg.Output)
}

func runFromDSN(arg RunArg) error {
	dxl, err := parser.From(arg.DSN, arg.Table...)
	if err != nil {
		return err
	}

	return run(dxl, arg.Mode, arg.Output)
}

var funcMap = map[Mode]func(dxl *spec.DXL, output string) error{
	SQL:  sql.Run,
	GORM: gorm.Run,
	XORM: xorm.Run,
	SQLX: sqlx.Run,
	BUN:  bun.Run,
}

func run(dxl *spec.DXL, mode Mode, output string) error {
	fn, ok := funcMap[mode]
	if !ok {
		return nil
	}
	return fn(dxl, output)
}
