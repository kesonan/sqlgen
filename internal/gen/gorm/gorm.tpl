package model

type {{UpperCamel .Name}} struct { {{range .Columns}}
    {{UpperCamel .Name}} {{.Go}} `gorm:"{{.Name}}" json:"{{LowerCamel .Name}}"`{{end}}
}