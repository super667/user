package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRpc zrpc.RpcClientConf

	MySql struct {
		DataSource string `json:"DataSource"`
	}

	Auth struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshExpire int64
	}
}
