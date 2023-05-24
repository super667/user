// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	strategyFieldNames          = builder.RawFieldNames(&Strategy{})
	strategyRows                = strings.Join(strategyFieldNames, ",")
	strategyRowsExpectAutoSet   = strings.Join(stringx.Remove(strategyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	strategyRowsWithPlaceHolder = strings.Join(stringx.Remove(strategyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	strategyModel interface {
		Insert(ctx context.Context, data *Strategy) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Strategy, error)
		Update(ctx context.Context, data *Strategy) error
		Delete(ctx context.Context, id int64) error
	}

	defaultStrategyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Strategy struct {
		Id              int64        `db:"id"`
		Subject         string       `db:"subject"`           // 权限赋予的主体
		SubjectName     string       `db:"subject_name"`      // 权限赋予的主体名称
		SubjectType     string       `db:"subject_type"`      // 主体类型
		SubjectTypeName string       `db:"subject_type_name"` // 主体类型名称
		Resource        string       `db:"resource"`          // 资源
		ResourceName    string       `db:"resource_name"`     // 资源名称
		Perm            string       `db:"perm"`              // 权限点
		PermName        string       `db:"perm_name"`         // 权限点名称
		CreateTime      time.Time    `db:"create_time"`
		UpdateTime      time.Time    `db:"update_time"`
		DeleteTime      sql.NullTime `db:"delete_time"`
	}
)

func newStrategyModel(conn sqlx.SqlConn) *defaultStrategyModel {
	return &defaultStrategyModel{
		conn:  conn,
		table: "`strategy`",
	}
}

func (m *defaultStrategyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultStrategyModel) FindOne(ctx context.Context, id int64) (*Strategy, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", strategyRows, m.table)
	var resp Strategy
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStrategyModel) Insert(ctx context.Context, data *Strategy) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, strategyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Subject, data.SubjectName, data.SubjectType, data.SubjectTypeName, data.Resource, data.ResourceName, data.Perm, data.PermName, data.DeleteTime)
	return ret, err
}

func (m *defaultStrategyModel) Update(ctx context.Context, data *Strategy) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strategyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Subject, data.SubjectName, data.SubjectType, data.SubjectTypeName, data.Resource, data.ResourceName, data.Perm, data.PermName, data.DeleteTime, data.Id)
	return err
}

func (m *defaultStrategyModel) tableName() string {
	return m.table
}