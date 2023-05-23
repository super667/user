package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		Count(ctx context.Context) (int64, error)
		FindManyByPage(ctx context.Context, page, pageSize int64) ([]*User, error)
		PartialUpdate(ctx context.Context, newData *User) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *defaultUserModel) Count(ctx context.Context) (int64, error) {
	var count int64
	rowBuilder := squirrel.Select("count(*)").From(m.tableName()).Where(squirrel.Eq{"delete_time": nil})
	query, _, err := rowBuilder.ToSql()
	if err != nil {
		return 0, err
	}
	err = m.conn.QueryRow(&count, query)
	if err != nil {
		return 0, err
	}
	return count, nil

}

func (m *defaultUserModel) FindManyByPage(ctx context.Context, page, pageSize int64) ([]*User, error) {
	rowBuilder := squirrel.Select(userRows).From(m.table)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	rowBuilder = rowBuilder.Offset(uint64((page - 1) * pageSize)).Limit(uint64(pageSize)).Where(squirrel.Eq{"delete_time": nil})
	query, _, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var userInfos []*User
	err = m.conn.QueryRowsCtx(ctx, &userInfos, query)
	if err != nil {
		return nil, err
	}
	return userInfos, nil
}

func (m *defaultUserModel) PartialUpdate(ctx context.Context, newData *User) error {
	rowBuilder := squirrel.Update(m.tableName())
	if newData.Number != "" {
		rowBuilder = rowBuilder.Set("number", newData.Number)
	}
	if newData.Name != "" {
		rowBuilder = rowBuilder.Set("name", newData.Name)
	}

	if newData.GenderCode != 0 {
		rowBuilder = rowBuilder.Set("gender_code", newData.GenderCode)
	}
	if newData.Age != 0 {
		rowBuilder = rowBuilder.Set("age", newData.Age)
	}
	if newData.DeptCode != "" {
		rowBuilder = rowBuilder.Set("dept_code", newData.DeptCode)
	}
	if newData.DeptName != "" {
		rowBuilder = rowBuilder.Set("dept_name", newData.DeptName)
	}
	if newData.ManagerCode != "" {
		rowBuilder = rowBuilder.Set("manager_code", newData.ManagerCode)
	}

	if newData.ManagerName != "" {
		rowBuilder = rowBuilder.Set("manager_name", newData.ManagerName)
	}

	if newData.Phone != "" {
		rowBuilder = rowBuilder.Set("phone", newData.Phone)
	}

	if newData.Email != "" {
		rowBuilder = rowBuilder.Set("email", newData.Email)
	}
	rowBuilder = rowBuilder.Where(squirrel.Eq{"id": newData.Id})

	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = m.conn.ExecCtx(ctx, query, args...)
	return err
}
