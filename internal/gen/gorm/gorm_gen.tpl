package model

import (
    "context"

    "gorm.io/gorm"
)

// {{UpperCamel .Name}}Model represents a {{.Name}} model.
type {{UpperCamel .Name}}Model struct {
    db gorm.DB
}

// TODO(sqlgen): Add your own customize code here.
func (m *{{UpperCamel .Name}}Model)Customize(ctx context.Context,args...any) {

}