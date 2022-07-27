package spec

import (
	"fmt"

	"github.com/pingcap/parser/mysql"
)

const defaultThirdDecimalPkg = "github.com/shopspring/decimal"

type typeKey struct {
	tp       byte
	signed   bool
	thirdPkg string
}

var typeMapper = map[typeKey]string{
	typeKey{tp: mysql.TypeTiny}:                   "int8",
	typeKey{tp: mysql.TypeTiny, signed: true}:     "uint8",
	typeKey{tp: mysql.TypeShort}:                  "int16",
	typeKey{tp: mysql.TypeShort, signed: true}:    "uint16",
	typeKey{tp: mysql.TypeLong}:                   "int32",
	typeKey{tp: mysql.TypeLong, signed: true}:     "uint32",
	typeKey{tp: mysql.TypeFloat}:                  "float64",
	typeKey{tp: mysql.TypeDouble}:                 "float64",
	typeKey{tp: mysql.TypeTimestamp}:              "time.Time",
	typeKey{tp: mysql.TypeLonglong}:               "int64",
	typeKey{tp: mysql.TypeLonglong, signed: true}: "uint64",
	typeKey{tp: mysql.TypeInt24}:                  "int32",
	typeKey{tp: mysql.TypeInt24, signed: true}:    "uint32",
	typeKey{tp: mysql.TypeDate}:                   "time.Time",
	typeKey{tp: mysql.TypeDuration}:               "time.Time",
	typeKey{tp: mysql.TypeDatetime}:               "time.Time",
	typeKey{tp: mysql.TypeYear}:                   "string",
	typeKey{tp: mysql.TypeVarchar}:                "string",
	typeKey{tp: mysql.TypeBit}:                    "byte",
	typeKey{tp: mysql.TypeJSON}:                   "string",
	typeKey{
		tp:       mysql.TypeNewDecimal,
		thirdPkg: defaultThirdDecimalPkg,
	}: "byte",
	typeKey{tp: mysql.TypeEnum}:       "string",
	typeKey{tp: mysql.TypeSet}:        "string",
	typeKey{tp: mysql.TypeTinyBlob}:   "string",
	typeKey{tp: mysql.TypeMediumBlob}: "string",
	typeKey{tp: mysql.TypeLongBlob}:   "string",
	typeKey{tp: mysql.TypeBlob}:       "string",
	typeKey{tp: mysql.TypeVarString}:  "string",
	typeKey{tp: mysql.TypeString}:     "string",
}

// Type is the type of the column.
type Type byte

// Go returns the Go type of the column.
func (c Column) Go(unsigned bool) (string, error) {
	var key = typeKey{tp: c.TP, signed: unsigned}
	if c.TP == mysql.TypeNewDecimal {
		key.thirdPkg = defaultThirdDecimalPkg
	}
	goType, ok := typeMapper[key]
	if !ok {
		return "", fmt.Errorf("unsupported type}: %v", c.TP)
	}

	return goType, nil
}
