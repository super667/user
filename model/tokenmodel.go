package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TokenModel = (*customTokenModel)(nil)

var cacheTokenTokenPrefix = "cache:token:token:"

type (
	// TokenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTokenModel.
	TokenModel interface {
		tokenModel
		FindOneByToken(ctx context.Context, token string) (*Token, error)
	}

	customTokenModel struct {
		*defaultTokenModel
	}
)

// NewTokenModel returns a model for the database table.
func NewTokenModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TokenModel {
	return &customTokenModel{
		defaultTokenModel: newTokenModel(conn, c, opts...),
	}
}

func (m *defaultTokenModel) SetKey(ctx context.Context, key, value string) error {
	return m.CachedConn.SetCacheCtx(ctx, key, value)
}

func (m *defaultTokenModel) FindOneByToken(ctx context.Context, token string) (*Token, error) {
	tokenKey := fmt.Sprintf("%s%v", cacheTokenTokenPrefix, token)
	rowBuilder := squirrel.Select(tokenRows).From(m.table).
		Where(squirrel.Eq{"delete_time": nil}).
		Where(squirrel.Eq{"token": token}).
		Limit(1)

	query, args, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp Token
	err = m.QueryRowCtx(ctx, &resp, tokenKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, query, args...)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
