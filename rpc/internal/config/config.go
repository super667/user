package config

import (
	"github.com/super667/user/common/ldap"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	MySql struct {
		DataSource string `json:"DataSource"`
	}

	CacheRedis cache.CacheConf

	Salt string

	Ldap ldap.Config

	JWT struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshExpire int64
	}
}
