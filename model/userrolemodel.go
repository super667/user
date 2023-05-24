package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserRoleModel = (*customUserRoleModel)(nil)

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleModel.
	UserRoleModel interface {
		userRoleModel
		FindManyByPage(ctx context.Context, page, pageSize int64) ([]*UserRole, error)
		Count(ctx context.Context) (int64, error)
		PartialUpdate(ctx context.Context, newData *UserRole) error
	}

	customUserRoleModel struct {
		*defaultUserRoleModel
	}
)

// NewUserRoleModel returns a model for the database table.
func NewUserRoleModel(conn sqlx.SqlConn) UserRoleModel {
	return &customUserRoleModel{
		defaultUserRoleModel: newUserRoleModel(conn),
	}
}

func (m *defaultUserRoleModel) FindManyByPage(ctx context.Context, page, pageSize int64) ([]*UserRole, error) {
	rowBuilder := squirrel.Select(userRoleRows).From(m.table)
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
	var Infos []*UserRole
	err = m.conn.QueryRowsCtx(ctx, &Infos, query)
	if err != nil {
		return nil, err
	}
	return Infos, nil
}

func (m *defaultUserRoleModel) Count(ctx context.Context) (int64, error) {
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

func (m *defaultUserRoleModel) PartialUpdate(ctx context.Context, newData *UserRole) error {
	rowBuilder := squirrel.Update(m.tableName())
	if newData.UserId != 0 {
		rowBuilder = rowBuilder.Set("number", newData.UserId)
	}
	if newData.RoleId != 0 {
		rowBuilder = rowBuilder.Set("name", newData.RoleId)
	}

	rowBuilder = rowBuilder.Where(squirrel.Eq{"id": newData.Id})

	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = m.conn.ExecCtx(ctx, query, args...)
	return err
}