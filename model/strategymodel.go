package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StrategyModel = (*customStrategyModel)(nil)

type (
	// StrategyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStrategyModel.
	StrategyModel interface {
		strategyModel
		FindManyByPage(ctx context.Context, page, pageSize int64) ([]*Strategy, error)
		Count(ctx context.Context) (int64, error)
		PartialUpdate(ctx context.Context, newData *Strategy) error
	}

	customStrategyModel struct {
		*defaultStrategyModel
	}
)

// NewStrategyModel returns a model for the database table.
func NewStrategyModel(conn sqlx.SqlConn) StrategyModel {
	return &customStrategyModel{
		defaultStrategyModel: newStrategyModel(conn),
	}
}

func (m *defaultStrategyModel) FindManyByPage(ctx context.Context, page, pageSize int64) ([]*Strategy, error) {
	rowBuilder := squirrel.Select(strategyRows).From(m.table)
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
	var Infos []*Strategy
	err = m.conn.QueryRowsCtx(ctx, &Infos, query)
	if err != nil {
		return nil, err
	}
	return Infos, nil
}

func (m *defaultStrategyModel) Count(ctx context.Context) (int64, error) {
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

func (m *defaultStrategyModel) PartialUpdate(ctx context.Context, newData *Strategy) error {
	rowBuilder := squirrel.Update(m.tableName())
	if newData.Subject != "" {
		rowBuilder = rowBuilder.Set("number", newData.Subject)
	}
	if newData.SubjectName != "" {
		rowBuilder = rowBuilder.Set("name", newData.SubjectName)
	}

	if newData.SubjectType != "" {
		rowBuilder = rowBuilder.Set("gender_code", newData.SubjectType)
	}
	if newData.SubjectTypeName != "" {
		rowBuilder = rowBuilder.Set("age", newData.SubjectTypeName)
	}
	if newData.Resource != "" {
		rowBuilder = rowBuilder.Set("dept_code", newData.Resource)
	}
	if newData.ResourceName != "" {
		rowBuilder = rowBuilder.Set("dept_name", newData.ResourceName)
	}
	if newData.Perm != "" {
		rowBuilder = rowBuilder.Set("manager_code", newData.Perm)
	}

	if newData.PermName != "" {
		rowBuilder = rowBuilder.Set("manager_name", newData.PermName)
	}

	rowBuilder = rowBuilder.Where(squirrel.Eq{"id": newData.Id})

	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = m.conn.ExecCtx(ctx, query, args...)
	return err
}

func (m *defaultStrategyModel) TranBatchAddStrategy(ctx context.Context, Strategys []*Strategy) (err error) {
	fn := func(ctx context.Context, session sqlx.Session) error {
		rowBuilder := squirrel.Insert(m.tableName())
		for _, s := range Strategys {
			rowBuilder = rowBuilder.SetMap(
				squirrel.Eq{
					"subject":           s.Subject,
					"subject_name":      s.SubjectName,
					"subject_type":      s.SubjectType,
					"subject_type_name": s.SubjectTypeName,
					"resource":          s.Resource,
					"resource_name":     s.ResourceName,
					"perm":              s.Perm,
					"perm_name":         s.PermName,
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
