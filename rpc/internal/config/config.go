package config

import (
	"github.com/super667/user/common/ldap"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	MySql struct {
		DataSource string `json:"DataSource"`
	}

	Salt string

	Ldap ldap.Config
}
