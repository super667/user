package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	RoleModel interface {
		roleModel
		FindManyByPage(ctx context.Context, page, pageSize int64) ([]*Role, error)
		Count(ctx context.Context) (int64, error)
		PartialUpdate(ctx context.Context, newData *Role) error
	}

	customRoleModel struct {
		*defaultRoleModel
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(conn sqlx.SqlConn) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(conn),
	}
}

func (m *defaultRoleModel) FindManyByPage(ctx context.Context, page, pageSize int64) ([]*Role, error) {
	rowBuilder := squirrel.Select(roleRows).From(m.table)
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
	var Infos []*Role
	err = m.conn.QueryRowsCtx(ctx, &Infos, query)
	if err != nil {
		return nil, err
	}
	return Infos, nil
}

func (m *defaultRoleModel) Count(ctx context.Context) (int64, error) {
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

func (m *defaultRoleModel) PartialUpdate(ctx context.Context, newData *Role) error {
	rowBuilder := squirrel.Update(m.tableName())
	if newData.Name != "" {
		rowBuilder = rowBuilder.Set("name", newData.Name)
	}
	if newData.Desc != "" {
		rowBuilder = rowBuilder.Set("desc", newData.Desc)
	}

	rowBuilder = rowBuilder.Where(squirrel.Eq{"id": newData.Id})

	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = m.conn.ExecCtx(ctx, query, args...)
	return err
}