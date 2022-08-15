package sql

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	model "github.com/anqiansong/sqlgen/example/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	um  *model.UserModel
	ctx = context.TODO()
	db  *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("mysql", "root:mysqlpw@tcp(127.0.0.1:55000)/test?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	um = model.NewUserModel(db, getScanner())
	m.Run()
}

func mustInitDB(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = tx.ExecContext(ctx, `truncate table user`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = tx.ExecContext(ctx, `alter table user auto_increment=1`)
	if err != nil {
		log.Fatalln(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestCreate(t *testing.T) {
	t.Run("emptyData", initAndRun(func(t *testing.T) {
		err := um.Create(ctx)
		assert.Contains(t, err.Error(), "empty")
	}))

	t.Run("createOne", initAndRun(func(t *testing.T) {
		mockUser := mustMockUser()
		err := um.Create(ctx, mockUser)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), mockUser.Id)
	}))
	t.Run("createMultiple", initAndRun(func(t *testing.T) {
		var list []*model.User
		for i := 1; i <= 5; i++ {
			list = append(list, mustMockUser())
		}
		err := um.Create(ctx, list...)
		assert.NoError(t, err)
		for idx, item := range list {
			assert.Equal(t, uint64(idx+1), item.Id)
		}
	}))
}

func TestFindOne(t *testing.T) {
	t.Run("noRows", initAndRun(func(t *testing.T) {
		_, err := um.FindOne(ctx, model.FindOneWhereParameter{IdEqual: 1})
		assert.ErrorIs(t, err, sql.ErrNoRows)
	}))

	t.Run("findOne", initAndRun(func(t *testing.T) {
		mockUser := mustMockUser()
		err := um.Create(ctx, mockUser)
		assert.NoError(t, err)
		actual, err := um.FindOne(ctx, model.FindOneWhereParameter{IdEqual: mockUser.Id})
		assert.NoError(t, err)
		assertUserEqual(t, mockUser, actual)
	}))

}

func assertUserEqual(t *testing.T, expected, actual *model.User) {
	now := time.Now()
	expected.CreateAt = now
	expected.UpdateAt = now
	actual.CreateAt = now
	actual.UpdateAt = now
	assert.Equal(t, *expected, *actual)
}

func initAndRun(f func(t *testing.T)) func(t *testing.T) {
	mustInitDB(db)
	return func(t *testing.T) {
		f(t)
	}
}
