package spec

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/anqiansong/sqlgen/internal/buffer"
	"github.com/anqiansong/sqlgen/internal/parameter"
	"github.com/anqiansong/sqlgen/internal/set"
)

func (l *Limit) IsValid() bool {
	if l == nil {
		return false
	}
	return l.Count > 0
}

// SQL returns the clause condition strings.
func (l *Limit) SQL() (string, error) {
	sql, _, err := l.marshal()
	return sql, err
}

// ParameterStructure returns the parameter type structure.
func (l *Limit) ParameterStructure() (string, error) {
	_, parameters, err := l.marshal()
	if err != nil {
		return "", err
	}

	var writer = buffer.New()
	writer.Write(`// %s is a limit parameter structure.`, l.ParameterStructureName())
	writer.Write(`type %s struct {`, l.ParameterStructureName())
	for _, v := range parameters {
		writer.Write("%s %s", v.Column, v.Type)
	}

	writer.Write(`}`)

	return writer.String(), nil
}

// ParameterThirdImports returns the third package imports.
func (l *Limit) ParameterThirdImports() (string, error) {
	_, parameters, err := l.marshal()
	if err != nil {
		return "", err
	}
	var thirdPkgSet = set.From()
	for _, v := range parameters {
		if len(v.ThirdPkg) == 0 {
			continue
		}
		thirdPkgSet.Add(v.ThirdPkg)
	}

	return strings.Join(thirdPkgSet.String(), "\n"), nil
}

// Parameters returns the parameter variables.
func (l *Limit) Parameters(pkg string) (string, error) {
	_, parameters, err := l.marshal()
	if err != nil {
		return "", err
	}
	var list []string
	for _, v := range parameters {
		list = append(list, fmt.Sprintf("%s.%s", pkg, v.Column))
	}

	return strings.Join(list, ", "), nil
}

// ParameterStructureName returns the parameter structure name.
func (l *Limit) ParameterStructureName() string {
	if !l.IsValid() {
		return ""
	}

	return strcase.ToCamel(fmt.Sprintf("%sLimitParameter", l.Comment.FuncName))
}

func (l *Limit) marshal() (sql string, parameters parameter.Parameters, err error) {
	parameters = parameter.Empty
	if l == nil {
		return
	}

	sql = fmt.Sprintf("limit %d", l.Count)
	parameters = append(parameters, NewParameter("count", "uint", ""))
	if l.Offset > 0 {
		sql = fmt.Sprintf("limit %d, %d", l.Offset, l.Count)
		parameters = append(parameters, NewParameter("offset", "uint", ""))
	}

	return
}
