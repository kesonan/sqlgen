// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

// UserModel represents a user model.
type UserModel struct {
	db bun.IDB
}

// User represents a user struct data.
type User struct {
	bun.BaseModel `bun:"table:user"`
	Id            uint64    `bun:"id,pk,autoincrement;" json:"id"`
	Name          string    `bun:"name" json:"name"`
	Password      string    `bun:"password" json:"password"`
	Mobile        string    `bun:"mobile" json:"mobile"`
	Gender        string    `bun:"gender" json:"gender"`
	Nickname      string    `bun:"nickname" json:"nickname"`
	Type          int8      `bun:"type" json:"type"`
	CreateTime    time.Time `bun:"create_time" json:"createTime"`
	UpdateTime    time.Time `bun:"update_time" json:"updateTime"`
}

// DeleteWhereParameter is a where parameter structure.
type DeleteWhereParameter struct {
	IdEqual uint64
}

// DeleteByNameWhereParameter is a where parameter structure.
type DeleteByNameWhereParameter struct {
	NameEqual string
}

// DeleteByNameAndMobileWhereParameter is a where parameter structure.
type DeleteByNameAndMobileWhereParameter struct {
	NameEqual   string
	MobileEqual string
}

// DeleteOrderByIDWhereParameter is a where parameter structure.
type DeleteOrderByIDWhereParameter struct {
	IdEqual uint64
}

// DeleteOrderByIDLimitWhereParameter is a where parameter structure.
type DeleteOrderByIDLimitWhereParameter struct {
	IdEqual uint64
}

// NewUserModel creates a new user model.
func NewUserModel(db bun.IDB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// Create creates  user data.
func (m *UserModel) Create(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	var list []User
	for _, v := range data {
		list = append(list, *v)
	}

	_, err := m.db.NewInsert().Model(&list).Exec(ctx)
	return err
}

// Delete is generated from sql:
// delete from user where id = ?;
func (m *UserModel) Delete(ctx context.Context, where DeleteWhereParameter) error {
	var db = m.db.NewDelete()
	db.Model(&User{})
	db.Where(`id = ?`, where.IdEqual)
	_, err := db.Exec(ctx)
	return err
}

// DeleteByName is generated from sql:
// delete from user where name = ?;
func (m *UserModel) DeleteByName(ctx context.Context, where DeleteByNameWhereParameter) error {
	var db = m.db.NewDelete()
	db.Model(&User{})
	db.Where(`name = ?`, where.NameEqual)
	_, err := db.Exec(ctx)
	return err
}

// DeleteByNameAndMobile is generated from sql:
// delete from user where name = ? and mobile = ?;
func (m *UserModel) DeleteByNameAndMobile(ctx context.Context, where DeleteByNameAndMobileWhereParameter) error {
	var db = m.db.NewDelete()
	db.Model(&User{})
	db.Where(`name = ? AND mobile = ?`, where.NameEqual, where.MobileEqual)
	_, err := db.Exec(ctx)
	return err
}

// DeleteOrderByID is generated from sql:
// delete from user where id = ? order by id desc;
func (m *UserModel) DeleteOrderByID(ctx context.Context, where DeleteOrderByIDWhereParameter) error {
	var db = m.db.NewDelete()
	db.Model(&User{})
	db.Where(`id = ?`, where.IdEqual)
	_, err := db.Exec(ctx)
	return err
}

// DeleteOrderByIDLimit is generated from sql:
// delete from user where id = ? order by id desc limit 10;
func (m *UserModel) DeleteOrderByIDLimit(ctx context.Context, where DeleteOrderByIDLimitWhereParameter) error {
	var db = m.db.NewDelete()
	db.Model(&User{})
	db.Where(`id = ?`, where.IdEqual)
	_, err := db.Exec(ctx)
	return err
}
