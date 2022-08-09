// Code generated by sqlgen. DO NOT EDIT!

package model

import (
    "context"
    "database/sql"
    "fmt"
    "time"


    "github.com/shopspring/decimal"
)

// {{UpperCamel $.Table.Name}}Model represents a {{$.Table.Name}} model.
type {{UpperCamel $.Table.Name}}Model struct {
    scanner Scanner
    db sql.Conn
}

// {{UpperCamel $.Table.Name}} represents a {{$.Table.Name}} struct data.
type {{UpperCamel $.Table.Name}} struct { {{range $.Table.Columns}}
{{UpperCamel .Name}} {{.GoType}}{{end}}
}

{{range $stmt := .SelectStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Having.IsValid}}{{$stmt.Having.ParameterStructure "Having"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}{{$stmt.ReceiverStructure}}
{{end}}

{{range $stmt := .UpdateStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}
{{end}}

{{range $stmt := .DeleteStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}
{{end}}

func (m *{{UpperCamel $.Table.Name}}Model) SetScanner(scanner Scanner) {
    m.scanner = scanner
}

// Insert creates  {{$.Table.Name}} data.
func (m *{{UpperCamel $.Table.Name}}Model) Insert(ctx context.Context, data ...*{{UpperCamel $.Table.Name}}) error {
    if len(data)==0{
        return fmt.Errorf("data is empty")
    }

    stmt,err := m.db.PrepareContext(ctx, {{InsertSQL}})
    if err != nil{
        return err
    }
    defer stmt.Close()

    for _,v := range data{
        _,err = stmt.ExecContext(ctx, {{InsertSQLArgs "v"}})
        if err != nil{
            return err
        }
    }

    return nil
}