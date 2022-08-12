// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UserModel represents a user model.
type UserModel struct {
	db gorm.DB
}

// User represents a user struct data.
type User struct {
	Id         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Password   string    `gorm:"column:password" json:"password"`
	Mobile     string    `gorm:"column:mobile" json:"mobile"`
	Gender     string    `gorm:"column:gender" json:"gender"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	Type       int8      `gorm:"column:type" json:"type"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// UpdateWhereParameter is a where parameter structure.
type UpdateWhereParameter struct {
	IdEqual uint64
}

// UpdateByNameWhereParameter is a where parameter structure.
type UpdateByNameWhereParameter struct {
	NameEqual string
}

// UpdatePartWhereParameter is a where parameter structure.
type UpdatePartWhereParameter struct {
	IdEqual uint64
}

// UpdatePartByNameWhereParameter is a where parameter structure.
type UpdatePartByNameWhereParameter struct {
	NameEqual string
}

// UpdateNameLimitWhereParameter is a where parameter structure.
type UpdateNameLimitWhereParameter struct {
	IdGT uint64
}

// UpdateNameLimitLimitParameter is a limit parameter structure.
type UpdateNameLimitLimitParameter struct {
	Count int
}

// UpdateNameLimitOrderWhereParameter is a where parameter structure.
type UpdateNameLimitOrderWhereParameter struct {
	IdGT uint64
}

// UpdateNameLimitOrderLimitParameter is a limit parameter structure.
type UpdateNameLimitOrderLimitParameter struct {
	Count int
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (User) TableName() string {
	return "user"
}

// NewUserModel returns a new user model.
func NewUserModel(db gorm.DB) *UserModel {
	return &UserModel{db: db}
}

// Create creates  user data.
func (m *UserModel) Create(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	db := m.db.WithContext(ctx)
	var list []User
	for _, v := range data {
		list = append(list, *v)
	}

	return db.Create(&list).Error
}

// Update is generated from sql:
// update user set name = ?, password = ?, mobile = ?, gender = ?, nickname = ?, type = ?, create_time = ?, update_time = ? where id = ?;
func (m *UserModel) Update(ctx context.Context, data *User, where UpdateWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db.Model(&User{})
	db.Where(`id = ?`, where.IdEqual)
	db.Updates(map[string]interface{}{
		"name":        data.Name,
		"password":    data.Password,
		"mobile":      data.Mobile,
		"gender":      data.Gender,
		"nickname":    data.Nickname,
		"type":        data.Type,
		"create_time": data.CreateTime,
		"update_time": data.UpdateTime,
	})
	return db.Error
}

// UpdateByName is generated from sql:
// update user set password = ?, mobile = ?, gender = ?, nickname = ?, type = ?, create_time = ?, update_time = ? where name = ?;
func (m *UserModel) UpdateByName(ctx context.Context, data *User, where UpdateByNameWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db.Model(&User{})
	db.Where(`name = ?`, where.NameEqual)
	db.Updates(map[string]interface{}{
		"password":    data.Password,
		"mobile":      data.Mobile,
		"gender":      data.Gender,
		"nickname":    data.Nickname,
		"type":        data.Type,
		"create_time": data.CreateTime,
		"update_time": data.UpdateTime,
	})
	return db.Error
}

// UpdatePart is generated from sql:
// update user set name = ?, nickname = ? where id = ?;
func (m *UserModel) UpdatePart(ctx context.Context, data *User, where UpdatePartWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db.Model(&User{})
	db.Where(`id = ?`, where.IdEqual)
	db.Updates(map[string]interface{}{
		"name":     data.Name,
		"nickname": data.Nickname,
	})
	return db.Error
}

// UpdatePartByName is generated from sql:
// update user set name = ?, nickname = ? where name = ?;
func (m *UserModel) UpdatePartByName(ctx context.Context, data *User, where UpdatePartByNameWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db.Model(&User{})
	db.Where(`name = ?`, where.NameEqual)
	db.Updates(map[string]interface{}{
		"name":     data.Name,
		"nickname": data.Nickname,
	})
	return db.Error
}

// UpdateNameLimit is generated from sql:
// update user set name = ? where id > ? limit ?;
func (m *UserModel) UpdateNameLimit(ctx context.Context, data *User, where UpdateNameLimitWhereParameter, limit UpdateNameLimitLimitParameter) error {
	var db = m.db.WithContext(ctx)
	db.Model(&User{})
	db.Where(`id > ?`, where.IdGT)
	db.Limit(limit.Count)
	db.Updates(map[string]interface{}{
		"name": data.Name,
	})
	return db.Error
}

// UpdateNameLimitOrder is generated from sql:
// update user set name = ? where id > ? order by id desc limit ?;
func (m *UserModel) UpdateNameLimitOrder(ctx context.Context, data *User, where UpdateNameLimitOrderWhereParameter, limit UpdateNameLimitOrderLimitParameter) error {
	var db = m.db.WithContext(ctx)
	db.Model(&User{})
	db.Where(`id > ?`, where.IdGT)
	db.Order(`id desc`)
	db.Limit(limit.Count)
	db.Updates(map[string]interface{}{
		"name": data.Name,
	})
	return db.Error
}
