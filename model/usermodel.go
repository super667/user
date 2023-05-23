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
	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return 0, err
	}
	err = m.conn.QueryRow(&count, query, args)
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
	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var userInfos []*User
	err = m.conn.QueryRowsCtx(ctx, userInfos, query, args)
	if err != nil {
		return nil, err
	}
	return userInfos, nil
}
