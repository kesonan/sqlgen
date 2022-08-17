package spec

import (
	"database/sql"
	"fmt"

	"github.com/anqiansong/sqlgen/internal/parameter"
	"github.com/pingcap/parser/mysql"
)

const (
	// TypeNullLongLong is a type extension for mysql.TypeLongLong.
	TypeNullLongLong byte = 0xf0
	// TypeNullDecimal is a type extension for mysql.TypeDecimal.
	TypeNullDecimal byte = 0xf1
	// TypeNullString is a type extension for mysql.TypeString.
	TypeNullString byte = 0xf2
)

const defaultThirdDecimalPkg = "github.com/shopspring/decimal"

type typeKey struct {
	tp            byte
	signed        bool
	thirdPkg      string
	aggregateCall bool
	sql.NullFloat64
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
	}: "decimal.Decimal",
	typeKey{
		tp:       TypeNullDecimal,
		thirdPkg: defaultThirdDecimalPkg,
	}: "decimal.NullDecimal",
	typeKey{tp: mysql.TypeEnum}:       "string",
	typeKey{tp: mysql.TypeSet}:        "string",
	typeKey{tp: mysql.TypeTinyBlob}:   "string",
	typeKey{tp: mysql.TypeMediumBlob}: "string",
	typeKey{tp: mysql.TypeLongBlob}:   "string",
	typeKey{tp: mysql.TypeBlob}:       "string",
	typeKey{tp: mysql.TypeVarString}:  "string",
	typeKey{tp: mysql.TypeString}:     "string",
	typeKey{tp: TypeNullString}:       "sql.NullString",

	// aggregate functions
	typeKey{tp: mysql.TypeTiny, aggregateCall: true}:     "sql.NullInt16",
	typeKey{tp: mysql.TypeShort, aggregateCall: true}:    "sql.NullInt16",
	typeKey{tp: mysql.TypeLong, aggregateCall: true}:     "sql.NullInt32",
	typeKey{tp: mysql.TypeFloat, aggregateCall: true}:    "sql.NullInt32",
	typeKey{tp: mysql.TypeDouble, aggregateCall: true}:   "sql.NullFloat64",
	typeKey{tp: mysql.TypeLonglong, aggregateCall: true}: "sql.NullInt64",
	typeKey{tp: mysql.TypeInt24, aggregateCall: true}:    "sql.NullInt32",
	typeKey{tp: mysql.TypeYear, aggregateCall: true}:     "sql.NullString",
	typeKey{tp: mysql.TypeVarchar, aggregateCall: true}:  "sql.NullString",
	typeKey{tp: mysql.TypeBit, aggregateCall: true}:      "sql.NullInt16",
	typeKey{tp: mysql.TypeJSON, aggregateCall: true}:     "sql.NullString",
	typeKey{
		tp:            mysql.TypeNewDecimal,
		thirdPkg:      defaultThirdDecimalPkg,
		aggregateCall: true,
	}: "decimal.NullDecimal",
	typeKey{
		tp:            TypeNullDecimal,
		thirdPkg:      defaultThirdDecimalPkg,
		aggregateCall: true,
	}: "decimal.NullDecimal",
	typeKey{tp: mysql.TypeEnum, aggregateCall: true}:       "sql.NullString",
	typeKey{tp: mysql.TypeSet, aggregateCall: true}:        "sql.NullString",
	typeKey{tp: mysql.TypeTinyBlob, aggregateCall: true}:   "sql.NullString",
	typeKey{tp: mysql.TypeMediumBlob, aggregateCall: true}: "sql.NullString",
	typeKey{tp: mysql.TypeLongBlob, aggregateCall: true}:   "sql.NullString",
	typeKey{tp: mysql.TypeBlob, aggregateCall: true}:       "sql.NullString",
	typeKey{tp: mysql.TypeVarString, aggregateCall: true}:  "sql.NullString",
	typeKey{tp: mysql.TypeString, aggregateCall: true}:     "sql.NullString",
	typeKey{tp: mysql.TypeString, aggregateCall: true}:     "sql.NullString",
	typeKey{tp: TypeNullLongLong}:                          "sql.NullInt64",
	typeKey{tp: TypeNullDecimal}:                           "decimal.NullDecimal",
	typeKey{tp: TypeNullString}:                            "sql.NullString",
	typeKey{tp: TypeNullLongLong, aggregateCall: true}:     "sql.NullInt64",
	typeKey{tp: TypeNullDecimal, aggregateCall: true}:      "decimal.NullDecimal",
	typeKey{tp: TypeNullString, aggregateCall: true}:       "sql.NullString",
}

// Type is the type of the column.
type Type byte

// DataType returns the Go type, third-package of the column.
func (c Column) DataType() (parameter.Parameter, error) {
	var key = typeKey{tp: c.TP, signed: c.Unsigned, aggregateCall: c.AggregateCall}
	if c.AggregateCall {
		key = typeKey{tp: c.TP, aggregateCall: c.AggregateCall}
	}
	if c.TP == mysql.TypeNewDecimal {
		key.thirdPkg = defaultThirdDecimalPkg
	}

	goType, ok := typeMapper[key]
	if !ok {
		return parameter.Parameter{}, fmt.Errorf("unsupported type: %v", c.TP)
	}

	return NewParameter(c.Name, goType, key.thirdPkg), nil
}

// GoType returns the Go type of the column.
func (c Column) GoType() (string, error) {
	p, err := c.DataType()
	return p.Type, err
}

func isNullType(tp byte) bool {
	return tp >= TypeNullLongLong && tp <= TypeNullString
}
