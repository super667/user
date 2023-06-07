package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindOneByName(ctx context.Context, name string) (*User, error)
		FindOneByAccount(ctx context.Context, account string) (*User, error)
		Count(ctx context.Context, search string) (int64, error)
		FindManyByPage(ctx context.Context, search string, page, pageSize int64) ([]*User, error)
		PartialUpdate(ctx context.Context, newData *User) error
		GetOne(ctx context.Context, phone string) (*User, error)
		Exist(ctx context.Context, userName string) (bool, error)
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

func (m *defaultUserModel) Exist(ctx context.Context, userName string) (bool, error) {
	user, err := m.FindOneByUserName(ctx, userName)
	if err == ErrNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}
	return false, nil
}

func (m *defaultUserModel) FindOneByUserName(ctx context.Context, name string) (*User, error) {
	rowBuilder := squirrel.Select(userRows).From(m.table)
	rowBuilder = rowBuilder.Where(squirrel.Eq{"user_name": name}).
		Where(squirrel.Eq{"delete_time": nil})
	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp User
	err = m.conn.QueryRowCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserModel) Count(ctx context.Context, search string) (int64, error) {
	var count int64
	rowBuilder := squirrel.Select("count(*)").From(m.tableName()).Where(squirrel.Eq{"delete_time": nil})
	if strings.TrimSpace(search) != "" {
		rowBuilder = rowBuilder.Where(squirrel.Like{"nick_name": "%" + search + "%"})
	}
	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return 0, err
	}
	err = m.conn.QueryRow(&count, query, args...)
	if err != nil {
		return 0, err
	}
	return count, nil

}

func (m *defaultUserModel) FindManyByPage(ctx context.Context, search string, page, pageSize int64) ([]*User, error) {
	rowBuilder := squirrel.Select(userRows).From(m.table)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if strings.TrimSpace(search) != "" {
		rowBuilder = rowBuilder.Where(squirrel.Like{"nick_name": "%" + search + "%"})
	}
	rowBuilder = rowBuilder.Offset(uint64((page - 1) * pageSize)).Limit(uint64(pageSize)).Where(squirrel.Eq{"delete_time": nil})
	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var userInfos []*User
	err = m.conn.QueryRowsCtx(ctx, &userInfos, query, args...)
	if err != nil {
		return nil, err
	}
	return userInfos, nil
}

func (m *defaultUserModel) GetOne(ctx context.Context, account string) (*User, error) {
	rowBuilder := squirrel.Select(userRows).From(m.table)
	rowBuilder = rowBuilder.Where(squirrel.Or{squirrel.Eq{"phone": account}, squirrel.Eq{"user_name": account}}).
		Where(squirrel.Eq{"delete_time": nil})
	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp User
	err = m.conn.QueryRowCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserModel) PartialUpdate(ctx context.Context, newData *User) error {
	rowBuilder := squirrel.Update(m.tableName())
	if newData.Number != "" {
		rowBuilder = rowBuilder.Set("number", newData.Number)
	}
	if newData.UserName != "" {
		rowBuilder = rowBuilder.Set("user_name", newData.UserName)
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

func (m *defaultUserModel) TranBatchAddUser(ctx context.Context, userList []*User) (err error) {
	fn := func(ctx context.Context, session sqlx.Session) error {
		rowBuilder := squirrel.Insert(m.tableName())
		for _, user := range userList {
			rowBuilder = rowBuilder.SetMap(
				squirrel.Eq{
					"user_name": user.UserName,
					"nick_name": user.NickName,
					"number":    user.Number,
				},
			)
			query, args, err := rowBuilder.ToSql()
			if err != nil {
				return err
			}
			_, err = session.ExecCtx(ctx, query, args...)
			if err != nil {
				return err
			}
			return nil
		}
		return nil
	}
	err = m.conn.TransactCtx(ctx, fn)
	return err
}

func (m *defaultUserModel) FindOneByName(ctx context.Context, name string) (*User, error) {
	return nil, nil
}
func (m *defaultUserModel) FindOneByAccount(ctx context.Context, account string) (*User, error) {
	return nil, nil
}
