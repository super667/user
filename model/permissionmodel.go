package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PermissionModel = (*customPermissionModel)(nil)

type (
	// PermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPermissionModel.
	PermissionModel interface {
		permissionModel
		FindManyByPage(ctx context.Context, page, pageSize int64) ([]*Permission, error)
		Count(ctx context.Context) (int64, error)
		PartialUpdate(ctx context.Context, newData *Permission) error
	}

	customPermissionModel struct {
		*defaultPermissionModel
	}
)

// NewPermissionModel returns a model for the database table.
func NewPermissionModel(conn sqlx.SqlConn) PermissionModel {
	return &customPermissionModel{
		defaultPermissionModel: newPermissionModel(conn),
	}
}

func (m *defaultPermissionModel) FindManyByPage(ctx context.Context, page, pageSize int64) ([]*Permission, error) {
	rowBuilder := squirrel.Select(permissionRows).From(m.table)
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
	var Infos []*Permission
	err = m.conn.QueryRowsCtx(ctx, &Infos, query)
	if err != nil {
		return nil, err
	}
	return Infos, nil
}

func (m *defaultPermissionModel) Count(ctx context.Context) (int64, error) {
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

func (m *defaultPermissionModel) PartialUpdate(ctx context.Context, newData *Permission) error {
	rowBuilder := squirrel.Update(m.tableName())
	if newData.Resource != "" {
		rowBuilder = rowBuilder.Set("resource", newData.Resource)
	}
	if newData.Perm != "" {
		rowBuilder = rowBuilder.Set("perm", newData.Perm)
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
