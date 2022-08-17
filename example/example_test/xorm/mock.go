package sql

import (
	"time"

	model "github.com/anqiansong/sqlgen/example/xorm"
	uuid "github.com/satori/go.uuid"
)

func mustMockUser() *model.User {
	uid := uuid.NewV4().String()
	now := time.Now()
	return &model.User{
		Name:     uid,
		Password: "bar",
		Mobile:   uid,
		Gender:   "male",
		Nickname: "test",
		Type:     1,
		CreateAt: now,
		UpdateAt: now,
	}
}
