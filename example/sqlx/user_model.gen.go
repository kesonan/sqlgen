// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"xorm.io/builder"
)

// UserModel represents a user model.
type UserModel struct {
	db *sqlx.DB
}

// User represents a user struct data.
type User struct {
	Id       uint64    `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	Password string    `db:"password" json:"password"`
	Mobile   string    `db:"mobile" json:"mobile"`
	Gender   string    `db:"gender" json:"gender"`
	Nickname string    `db:"nickname" json:"nickname"`
	Type     int8      `db:"type" json:"type"`
	CreateAt time.Time `db:"create_at" json:"createAt"`
	UpdateAt time.Time `db:"update_at" json:"updateAt"`
}

// FindOneWhereParameter is a where parameter structure.
type FindOneWhereParameter struct {
	IdEqual uint64
}

// FindOneByNameWhereParameter is a where parameter structure.
type FindOneByNameWhereParameter struct {
	NameEqual string
}

// FindOneGroupByNameWhereParameter is a where parameter structure.
type FindOneGroupByNameWhereParameter struct {
	NameEqual string
}

// FindOneGroupByNameHavingNameWhereParameter is a where parameter structure.
type FindOneGroupByNameHavingNameWhereParameter struct {
	NameEqual string
}

// FindOneGroupByNameHavingNameHavingParameter is a having parameter structure.
type FindOneGroupByNameHavingNameHavingParameter struct {
	NameEqual string
}

// FindLimitWhereParameter is a where parameter structure.
type FindLimitWhereParameter struct {
	IdGT uint64
}

// FindLimitLimitParameter is a limit parameter structure.
type FindLimitLimitParameter struct {
	Count int
}

// FindLimitOffsetLimitParameter is a limit parameter structure.
type FindLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindGroupLimitOffsetWhereParameter is a where parameter structure.
type FindGroupLimitOffsetWhereParameter struct {
	IdGT uint64
}

// FindGroupLimitOffsetLimitParameter is a limit parameter structure.
type FindGroupLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindGroupHavingLimitOffsetWhereParameter is a where parameter structure.
type FindGroupHavingLimitOffsetWhereParameter struct {
	IdGT uint64
}

// FindGroupHavingLimitOffsetHavingParameter is a having parameter structure.
type FindGroupHavingLimitOffsetHavingParameter struct {
	IdGT uint64
}

// FindGroupHavingLimitOffsetLimitParameter is a limit parameter structure.
type FindGroupHavingLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindGroupHavingOrderAscLimitOffsetWhereParameter is a where parameter structure.
type FindGroupHavingOrderAscLimitOffsetWhereParameter struct {
	IdGT uint64
}

// FindGroupHavingOrderAscLimitOffsetHavingParameter is a having parameter structure.
type FindGroupHavingOrderAscLimitOffsetHavingParameter struct {
	IdGT uint64
}

// FindGroupHavingOrderAscLimitOffsetLimitParameter is a limit parameter structure.
type FindGroupHavingOrderAscLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindGroupHavingOrderDescLimitOffsetWhereParameter is a where parameter structure.
type FindGroupHavingOrderDescLimitOffsetWhereParameter struct {
	IdGT uint64
}

// FindGroupHavingOrderDescLimitOffsetHavingParameter is a having parameter structure.
type FindGroupHavingOrderDescLimitOffsetHavingParameter struct {
	IdGT uint64
}

// FindGroupHavingOrderDescLimitOffsetLimitParameter is a limit parameter structure.
type FindGroupHavingOrderDescLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindOnePartWhereParameter is a where parameter structure.
type FindOnePartWhereParameter struct {
	IdGT uint64
}

// FindAllCountResult is a find all count result.
type FindAllCountResult struct {
	CountID sql.NullInt64 `db:"countID" json:"countID"`
}

// FindAllCountWhereWhereParameter is a where parameter structure.
type FindAllCountWhereWhereParameter struct {
	IdGT uint64
}

// FindAllCountWhereResult is a find all count where result.
type FindAllCountWhereResult struct {
	CountID sql.NullInt64 `db:"countID" json:"countID"`
}

// FindMaxIDResult is a find max id result.
type FindMaxIDResult struct {
	MaxID sql.NullInt64 `db:"maxID" json:"maxID"`
}

// FindMinIDResult is a find min id result.
type FindMinIDResult struct {
	MinID sql.NullInt64 `db:"minID" json:"minID"`
}

// FindAvgIDResult is a find avg id result.
type FindAvgIDResult struct {
	AvgID sql.NullInt64 `db:"avgID" json:"avgID"`
}

// UpdateWhereParameter is a where parameter structure.
type UpdateWhereParameter struct {
	IdEqual uint64
}

// UpdateOrderByIdDescWhereParameter is a where parameter structure.
type UpdateOrderByIdDescWhereParameter struct {
	IdEqual uint64
}

// UpdateOrderByIdDescLimitCountWhereParameter is a where parameter structure.
type UpdateOrderByIdDescLimitCountWhereParameter struct {
	IdEqual uint64
}

// DeleteOneWhereParameter is a where parameter structure.
type DeleteOneWhereParameter struct {
	IdEqual uint64
}

// DeleteOneByNameWhereParameter is a where parameter structure.
type DeleteOneByNameWhereParameter struct {
	NameEqual string
}

// DeleteOneOrderByIDAscWhereParameter is a where parameter structure.
type DeleteOneOrderByIDAscWhereParameter struct {
	NameEqual string
}

// DeleteOneOrderByIDDescWhereParameter is a where parameter structure.
type DeleteOneOrderByIDDescWhereParameter struct {
	NameEqual string
}

// DeleteOneOrderByIDDescLimitCountWhereParameter is a where parameter structure.
type DeleteOneOrderByIDDescLimitCountWhereParameter struct {
	NameEqual string
}

// DeleteOneOrderByIDDescLimitCountLimitParameter is a limit parameter structure.
type DeleteOneOrderByIDDescLimitCountLimitParameter struct {
	Count int
}

// NewUserModel creates a new user model.
func NewUserModel(db *sqlx.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// Create creates  user data.
func (m *UserModel) Create(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	var stmt *sql.Stmt
	stmt, err := m.db.PrepareContext(ctx, "INSERT INTO user (`name`, `password`, `mobile`, `gender`, `nickname`, `type`, `create_at`, `update_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range data {
		result, err := stmt.ExecContext(ctx, v.Name, v.Password, v.Mobile, v.Gender, v.Nickname, v.Type, v.CreateAt, v.UpdateAt)
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		v.Id = uint64(id)
	}
	return nil
}

// FindOne is generated from sql:
// select * from `user` where `id` = ? limit 1;
func (m *UserModel) FindOne(ctx context.Context, where FindOneWhereParameter) (result *User, err error) {
	result = new(User)
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`id = ?`, where.IdEqual))
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindOneByName is generated from sql:
// select * from `user` where `name` = ? limit 1;
func (m *UserModel) FindOneByName(ctx context.Context, where FindOneByNameWhereParameter) (result *User, err error) {
	result = new(User)
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`name = ?`, where.NameEqual))
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindOneGroupByName is generated from sql:
// select * from `user` where `name` = ? group by name limit 1;
func (m *UserModel) FindOneGroupByName(ctx context.Context, where FindOneGroupByNameWhereParameter) (result *User, err error) {
	result = new(User)
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`name = ?`, where.NameEqual))
	b.GroupBy(`name`)
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindOneGroupByNameHavingName is generated from sql:
// select * from `user` where `name` = ? group by name having name = ? limit 1;
func (m *UserModel) FindOneGroupByNameHavingName(ctx context.Context, where FindOneGroupByNameHavingNameWhereParameter, having FindOneGroupByNameHavingNameHavingParameter) (result *User, err error) {
	result = new(User)
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`name = ?`, where.NameEqual))
	b.GroupBy(`name`)
	b.Having(fmt.Sprintf(`name = '%v'`, having.NameEqual))
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindAll is generated from sql:
// select * from `user`;
func (m *UserModel) FindAll(ctx context.Context) (result []*User, err error) {
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, &v)
	}

	return result, nil
}

// FindLimit is generated from sql:
// select * from `user` where id > ? limit ?;
func (m *UserModel) FindLimit(ctx context.Context, where FindLimitWhereParameter, limit FindLimitLimitParameter) (result []*User, err error) {
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`id > ?`, where.IdGT))
	b.Limit(limit.Count)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, &v)
	}

	return result, nil
}

// FindLimitOffset is generated from sql:
// select * from `user` limit ?, ?;
func (m *UserModel) FindLimitOffset(ctx context.Context, limit FindLimitOffsetLimitParameter) (result []*User, err error) {
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Limit(limit.Count, limit.Offset)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, &v)
	}

	return result, nil
}

// FindGroupLimitOffset is generated from sql:
// select * from `user` where id > ? group by name limit ?, ?;
func (m *UserModel) FindGroupLimitOffset(ctx context.Context, where FindGroupLimitOffsetWhereParameter, limit FindGroupLimitOffsetLimitParameter) (result []*User, err error) {
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`id > ?`, where.IdGT))
	b.GroupBy(`name`)
	b.Limit(limit.Count, limit.Offset)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, &v)
	}

	return result, nil
}

// FindGroupHavingLimitOffset is generated from sql:
// select * from `user` where id > ? group by name having id > ? limit ?, ?;
func (m *UserModel) FindGroupHavingLimitOffset(ctx context.Context, where FindGroupHavingLimitOffsetWhereParameter, having FindGroupHavingLimitOffsetHavingParameter, limit FindGroupHavingLimitOffsetLimitParameter) (result []*User, err error) {
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`id > ?`, where.IdGT))
	b.GroupBy(`name`)
	b.Having(fmt.Sprintf(`id > '%v'`, having.IdGT))
	b.Limit(limit.Count, limit.Offset)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, &v)
	}

	return result, nil
}

// FindGroupHavingOrderAscLimitOffset is generated from sql:
// select * from `user` where id > ? group by name having id > ? order by id limit ?, ?;
func (m *UserModel) FindGroupHavingOrderAscLimitOffset(ctx context.Context, where FindGroupHavingOrderAscLimitOffsetWhereParameter, having FindGroupHavingOrderAscLimitOffsetHavingParameter, limit FindGroupHavingOrderAscLimitOffsetLimitParameter) (result []*User, err error) {
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`id > ?`, where.IdGT))
	b.GroupBy(`name`)
	b.Having(fmt.Sprintf(`id > '%v'`, having.IdGT))
	b.OrderBy(`id`)
	b.Limit(limit.Count, limit.Offset)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, &v)
	}

	return result, nil
}

// FindGroupHavingOrderDescLimitOffset is generated from sql:
// select * from `user` where id > ? group by name having id > ? order by id desc limit ?, ?;
func (m *UserModel) FindGroupHavingOrderDescLimitOffset(ctx context.Context, where FindGroupHavingOrderDescLimitOffsetWhereParameter, having FindGroupHavingOrderDescLimitOffsetHavingParameter, limit FindGroupHavingOrderDescLimitOffsetLimitParameter) (result []*User, err error) {
	b := builder.MySQL()
	b.Select(`*`)
	b.From("`user`")
	b.Where(builder.Expr(`id > ?`, where.IdGT))
	b.GroupBy(`name`)
	b.Having(fmt.Sprintf(`id > '%v'`, having.IdGT))
	b.OrderBy(`id desc`)
	b.Limit(limit.Count, limit.Offset)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, &v)
	}

	return result, nil
}

// FindOnePart is generated from sql:
// select `name`, `password`, `mobile` from `user` where id > ? limit 1;
func (m *UserModel) FindOnePart(ctx context.Context, where FindOnePartWhereParameter) (result *User, err error) {
	result = new(User)
	b := builder.MySQL()
	b.Select(`name, password, mobile`)
	b.From("`user`")
	b.Where(builder.Expr(`id > ?`, where.IdGT))
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v User
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindAllCount is generated from sql:
// select count(id) AS countID from `user`;
func (m *UserModel) FindAllCount(ctx context.Context) (result *FindAllCountResult, err error) {
	result = new(FindAllCountResult)
	b := builder.MySQL()
	b.Select(`count(id) AS countID`)
	b.From("`user`")
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v FindAllCountResult
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindAllCountWhere is generated from sql:
// select count(id) AS countID from `user` where id > ?;
func (m *UserModel) FindAllCountWhere(ctx context.Context, where FindAllCountWhereWhereParameter) (result *FindAllCountWhereResult, err error) {
	result = new(FindAllCountWhereResult)
	b := builder.MySQL()
	b.Select(`count(id) AS countID`)
	b.From("`user`")
	b.Where(builder.Expr(`id > ?`, where.IdGT))
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v FindAllCountWhereResult
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindMaxID is generated from sql:
// select max(id) AS maxID from `user`;
func (m *UserModel) FindMaxID(ctx context.Context) (result *FindMaxIDResult, err error) {
	result = new(FindMaxIDResult)
	b := builder.MySQL()
	b.Select(`max(id) AS maxID`)
	b.From("`user`")
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v FindMaxIDResult
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindMinID is generated from sql:
// select min(id) AS minID from `user`;
func (m *UserModel) FindMinID(ctx context.Context) (result *FindMinIDResult, err error) {
	result = new(FindMinIDResult)
	b := builder.MySQL()
	b.Select(`min(id) AS minID`)
	b.From("`user`")
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v FindMinIDResult
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// FindAvgID is generated from sql:
// select avg(id) AS avgID from `user`;
func (m *UserModel) FindAvgID(ctx context.Context) (result *FindAvgIDResult, err error) {
	result = new(FindAvgIDResult)
	b := builder.MySQL()
	b.Select(`avg(id) AS avgID`)
	b.From("`user`")
	b.Limit(1)
	query, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	rows, err = m.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v FindAvgIDResult
		err = rows.StructScan(&v)
		if err != nil {
			return nil, err
		}
		result = &v
		break
	}

	return result, nil
}

// Update is generated from sql:
// update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ?;
func (m *UserModel) Update(ctx context.Context, data *User, where UpdateWhereParameter) error {
	b := builder.MySQL()
	b.Update(builder.Eq{
		"name":      data.Name,
		"password":  data.Password,
		"mobile":    data.Mobile,
		"gender":    data.Gender,
		"nickname":  data.Nickname,
		"type":      data.Type,
		"create_at": data.CreateAt,
		"update_at": data.UpdateAt,
	})
	b.From("`user`")
	b.Where(builder.Expr(`id = ?`, where.IdEqual))
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	_, err = m.db.ExecContext(ctx, query, args...)
	return err
}

// UpdateOrderByIdDesc is generated from sql:
// update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ? order by id desc;
func (m *UserModel) UpdateOrderByIdDesc(ctx context.Context, data *User, where UpdateOrderByIdDescWhereParameter) error {
	b := builder.MySQL()
	b.Update(builder.Eq{
		"name":      data.Name,
		"password":  data.Password,
		"mobile":    data.Mobile,
		"gender":    data.Gender,
		"nickname":  data.Nickname,
		"type":      data.Type,
		"create_at": data.CreateAt,
		"update_at": data.UpdateAt,
	})
	b.From("`user`")
	b.Where(builder.Expr(`id = ?`, where.IdEqual))
	b.OrderBy(`id desc`)
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	_, err = m.db.ExecContext(ctx, query, args...)
	return err
}

// UpdateOrderByIdDescLimitCount is generated from sql:
// update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ? order by id desc;
func (m *UserModel) UpdateOrderByIdDescLimitCount(ctx context.Context, data *User, where UpdateOrderByIdDescLimitCountWhereParameter) error {
	b := builder.MySQL()
	b.Update(builder.Eq{
		"name":      data.Name,
		"password":  data.Password,
		"mobile":    data.Mobile,
		"gender":    data.Gender,
		"nickname":  data.Nickname,
		"type":      data.Type,
		"create_at": data.CreateAt,
		"update_at": data.UpdateAt,
	})
	b.From("`user`")
	b.Where(builder.Expr(`id = ?`, where.IdEqual))
	b.OrderBy(`id desc`)
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	_, err = m.db.ExecContext(ctx, query, args...)
	return err
}

// DeleteOne is generated from sql:
// delete from `user` where `id` = ?;
func (m *UserModel) DeleteOne(ctx context.Context, where DeleteOneWhereParameter) error {
	b := builder.MySQL()
	b.Delete()
	b.From("`user`")
	b.Where(builder.Expr(`id = ?`, where.IdEqual))
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	_, err = m.db.ExecContext(ctx, query, args...)
	return err
}

// DeleteOneByName is generated from sql:
// delete from `user` where `name` = ?;
func (m *UserModel) DeleteOneByName(ctx context.Context, where DeleteOneByNameWhereParameter) error {
	b := builder.MySQL()
	b.Delete()
	b.From("`user`")
	b.Where(builder.Expr(`name = ?`, where.NameEqual))
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	_, err = m.db.ExecContext(ctx, query, args...)
	return err
}

// DeleteOneOrderByIDAsc is generated from sql:
// delete from `user` where `name` = ? order by id;
func (m *UserModel) DeleteOneOrderByIDAsc(ctx context.Context, where DeleteOneOrderByIDAscWhereParameter) error {
	b := builder.MySQL()
	b.Delete()
	b.From("`user`")
	b.Where(builder.Expr(`name = ?`, where.NameEqual))
	b.OrderBy(`id`)
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	_, err = m.db.ExecContext(ctx, query, args...)
	return err
}

// DeleteOneOrderByIDDesc is generated from sql:
// delete from `user` where `name` = ? order by id desc;
func (m *UserModel) DeleteOneOrderByIDDesc(ctx context.Context, where DeleteOneOrderByIDDescWhereParameter) error {
	b := builder.MySQL()
	b.Delete()
	b.From("`user`")
	b.Where(builder.Expr(`name = ?`, where.NameEqual))
	b.OrderBy(`id desc`)
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	_, err = m.db.ExecContext(ctx, query, args...)
	return err
}

// DeleteOneOrderByIDDescLimitCount is generated from sql:
// delete from `user` where `name` = ? order by id desc limit ?;
func (m *UserModel) DeleteOneOrderByIDDescLimitCount(ctx context.Context, where DeleteOneOrderByIDDescLimitCountWhereParameter, limit DeleteOneOrderByIDDescLimitCountLimitParameter) error {
	b := builder.MySQL()
	b.Delete()
	b.From("`user`")
	b.Where(builder.Expr(`name = ?`, where.NameEqual))
	b.OrderBy(`id desc`)
	b.Limit(limit.Count)
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	_, err = m.db.ExecContext(ctx, query, args...)
	return err
}
