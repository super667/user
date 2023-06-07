package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TokenModel = (*customTokenModel)(nil)

type (
	// TokenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTokenModel.
	TokenModel interface {
		tokenModel
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
