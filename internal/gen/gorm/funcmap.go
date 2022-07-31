package gorm

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/anqiansong/sqlgen/internal/spec"
)

type SelectStmt struct {
	// Action represents the db action.
	Action spec.Action
	// Columns represents the operation columns.
	Columns []string
	// Comment represents a sql comment.
	spec.Comment
	// Distinct represents the select distinct flag.
	Distinct bool
	// From represents the operation table name, do not support multiple tables.
	From string

	ParamTypeDeclaration string
	// where clause
	WhereClauseSQL string
	WhereParameter string

	// having clause
	HavingClauseSQL string
	HavingParameter string

	// group by

}

func clauseParamType(funcName string, table *spec.Table, clause *spec.Clause) (string, error) {
	tp, _, _, err := clauseCode(funcName, table, clause)
	return tp, err
}

func clauseCode(funcName string, table *spec.Table, clause *spec.Clause) (typeArg, whereClauseSQL, paramSql string, err error) {
	var argSet = map[Arg]int{}
	var argList []string
	var paramSqlList []string
	_, params, whereSql, err := getClause(table, clause)
	if err != nil {
		return "", "", "", err
	}

	for _, arg := range params {
		if v, ok := argSet[arg]; ok {
			name := fmt.Sprintf("%s%d", arg.Name, v+1)
			paramSqlList = append(paramSqlList, name)
			argList = append(argList, Arg{
				Name: name,
				Type: arg.Type,
			}.String())
			argSet[arg] = v + 1
			continue
		}
		argSet[arg] = 0
		paramSqlList = append(paramSqlList, arg.Name)
		argList = append(argList, arg.String())
	}

	typeArg = fmt.Sprintf(`type %sParam struct{
%s
}`, strcase.ToCamel(funcName), strings.Join(argList, "\n"))
	paramSql = strings.Join(paramSqlList, ",")
	whereClauseSQL = whereSql
	return
}

type Arg struct {
	Name string
	Type string
}

func (a Arg) String() string {
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}

func getClause(table *spec.Table, clause *spec.Clause) (columns []string, params []Arg, whereClauseSQL string, err error) {
	columns = []string{}
	if clause == nil {
		return
	}

	switch clause.OP {
	case spec.And, spec.Or:
		leftColumns, leftParams, leftWhereClause, err := getClause(table, clause.Left)
		if err != nil {
			return nil, nil, "", err
		}

		rightColumns, rightParams, rightWhereClause, err := getClause(table, clause.Right)
		if err != nil {
			return nil, nil, "", err
		}

		columns = append(columns, leftColumns...)
		columns = append(columns, rightColumns...)
		params = append(params, leftParams...)
		params = append(params, rightParams...)
		var lrw []string
		if len(leftWhereClause) > 0 {
			lrw = append(lrw, leftWhereClause)
		}
		if len(rightWhereClause) > 0 {
			lrw = append(lrw, rightWhereClause)
		}
		whereClauseSQL = strings.Join(lrw, " "+spec.Operator[clause.OP])
	case spec.EQ, spec.GE, spec.GT, spec.In, spec.LE, spec.LT, spec.Like, spec.NE, spec.NotIn, spec.NotLike:
		columns = append(columns, clause.Column)
		c, ok := table.GetColumnByName(clause.Column)
		if !ok {
			return nil, nil, "", fmt.Errorf("column %q is not in table %q", clause.Column, table.Name)
		}

		goType, err := c.Go()
		if err != nil {
			return nil, nil, "", err
		}

		params = append(params, Arg{
			Name: strcase.ToCamel(fmt.Sprintf("%s%sArg", clause.Column, spec.OpName[clause.OP])),
			Type: goType,
		})
		whereClauseSQL = fmt.Sprintf("%s %s ?", clause.Column, spec.Operator[clause.OP])
	case spec.Between, spec.NotBetween:
		c, ok := table.GetColumnByName(clause.Column)
		if !ok {
			return nil, nil, "", fmt.Errorf("column %q is not in table %q", clause.Column, table.Name)
		}

		goType, err := c.Go()
		if err != nil {
			return nil, nil, "", err
		}

		columns = append(columns, clause.Column)
		params = append(params, Arg{
			Name: strcase.ToCamel(fmt.Sprintf("%s%sArgStart", clause.Column, spec.OpName[clause.OP])),
			Type: goType,
		})
		params = append(params, Arg{
			Name: strcase.ToCamel(fmt.Sprintf("%s%sArgTerminal", clause.Column, spec.OpName[clause.OP])),
			Type: goType,
		})
		whereClauseSQL = fmt.Sprintf("%s %s ? AND ?", clause.Column, spec.Operator[clause.OP])
	default:
		// ignores other types
	}
	return
}
