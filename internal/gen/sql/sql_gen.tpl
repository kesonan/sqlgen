// Code generated by sqlgen. DO NOT EDIT!

package model

import (
    "context"
    "database/sql"
    "fmt"
    "time"

    "xorm.io/builder"
    "github.com/shopspring/decimal"
)

// {{UpperCamel $.Table.Name}}Model represents a {{$.Table.Name}} model.
type {{UpperCamel $.Table.Name}}Model struct {
    db      *sql.DB
    scanner Scanner
}

// {{UpperCamel $.Table.Name}} represents a {{$.Table.Name}} struct data.
type {{UpperCamel $.Table.Name}} struct { {{range $.Table.Columns}}
{{UpperCamel .Name}} {{.GoType}} `json:"{{LowerCamel .Name}}"`{{end}}
}
{{range $stmt := .SelectStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Having.IsValid}}{{$stmt.Having.ParameterStructure "Having"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}{{$stmt.ReceiverStructure "sql"}}
{{end}}

{{range $stmt := .UpdateStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}
{{end}}

{{range $stmt := .DeleteStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}
{{end}}


// New{{UpperCamel $.Table.Name}}Model creates a new {{$.Table.Name}} model.
func New{{UpperCamel $.Table.Name}}Model(db *sql.DB, scanner Scanner) *{{UpperCamel $.Table.Name}}Model {
    return &{{UpperCamel $.Table.Name}}Model{
        db: db,
        scanner: scanner,
    }
}

// Create creates  {{$.Table.Name}} data.
func (m *{{UpperCamel $.Table.Name}}Model) Create(ctx context.Context, data ...*{{UpperCamel $.Table.Name}}) error {
    if len(data) == 0 {
        return fmt.Errorf("data is empty")
    }

    var stmt *sql.Stmt
    stmt, err := m.db.PrepareContext(ctx, "INSERT INTO {{$.Table.Name}} ({{InsertSQL}}) VALUES ({{InsertQuotes}})")
    if err != nil {
        return err
    }
    defer stmt.Close()
    for _, v := range data {
        result, err := stmt.ExecContext(ctx, {{InsertValues "v"}})
        if err != nil {
            return err
        }

        id, err := result.LastInsertId()
        if err != nil {
            return err
        }

        {{range $.Table.Columns}}{{if IsPrimary .Name}}{{if .AutoIncrement}}v.{{UpperCamel .Name}} = {{.GoType}}(id){{end}}{{end}}{{end}}
    }
    return nil
}
{{range $stmt := .SelectStmt}}
// {{.FuncName}} is generated from sql:
// {{$stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Having.IsValid}}, having {{$stmt.Having.ParameterStructureName "Having"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}})(result {{if $stmt.Limit.One}}*{{$stmt.ReceiverName}}, {{else}}[]*{{$stmt.ReceiverName}}, {{end}} err error){ {{if $stmt.Limit.One}}
    result = new({{$stmt.ReceiverName}}){{end}}
    b := builder.MySQL()
    b.Select(`{{$stmt.SelectSQL}}`)
    b.From("`{{$.Table.Name}}`")
    {{if $stmt.Where.IsValid}}b.Where(builder.Expr({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}}))
    {{end }}{{if $stmt.GroupBy.IsValid}}b.GroupBy({{$stmt.GroupBy.SQL}})
    {{end}}{{if $stmt.Having.IsValid}}b.Having(fmt.Sprintf({{HavingSprintf $stmt.Having.SQL}}, {{$stmt.Having.Parameters "having"}}))
    {{end}}{{if $stmt.OrderBy.IsValid}}b.OrderBy({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}b.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}}{{if gt $stmt.Limit.Offset 0}}, {{$stmt.Limit.OffsetParameter "limit"}}{{end}})
    {{end}}query, args, err := b.ToSQL()
    if err != nil {
        return nil, err
    }

    {{if $stmt.Limit.One}}row := m.db.QueryRowContext(ctx, query, args...)
    if err = row.Err(); err != nil {
        return nil, err
    }
    err = m.scanner.ScanRow(row, result)
    return
    {{else}}var rows *sql.Rows
    rows, err = m.db.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer func() {
        err = rows.Close()
        if err != nil {
            result = nil
        }
    }()
    if err = m.scanner.ScanRows(rows, result); err != nil{
        return nil, err
    }
    return result, nil{{end}}
}
{{end}}

{{range $stmt := .UpdateStmt}}
// {{.FuncName}} is generated from sql:
// {{$stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context, data *{{UpperCamel $.Table.Name}}{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}}) error {
    b := builder.MySQL()
    b.Update(builder.Eq{
        {{range $name := $stmt.Columns}}"{{$name}}": data.{{UpperCamel $name}},
        {{end}}
    })
    b.From("`{{$.Table.Name}}`")
    {{if $stmt.Where.IsValid}}b.Where(builder.Expr({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}}))
    {{end}}{{if $stmt.OrderBy.IsValid}}b.OrderBy({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}b.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}}{{if gt $stmt.Limit.Offset 0}}, {{$stmt.Limit.OffsetParameter "limit"}}{{end}})
    {{end}}query, args, err := b.ToSQL()
    if err != nil {
        return err
    }
    _, err = m.db.ExecContext(ctx, query, args...)
    return err
}
{{end}}

{{range $stmt := .DeleteStmt}}
// {{.FuncName}} is generated from sql:
// {{$stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}}) error {
    b := builder.MySQL()
    b.Delete()
    b.From("`{{$.Table.Name}}`")
    {{if $stmt.Where.IsValid}}b.Where(builder.Expr({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}}))
    {{end}}{{if $stmt.OrderBy.IsValid}}b.OrderBy({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}b.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}}{{if gt $stmt.Limit.Offset 0}}, {{$stmt.Limit.OffsetParameter "limit"}}{{end}})
    {{end}}query, args, err := b.ToSQL()
    if err != nil {
        return err
    }
    _, err = m.db.ExecContext(ctx, query, args...)
    return err
}
{{end}}