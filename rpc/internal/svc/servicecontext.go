package svc

import (
	casbinsqlx "github.com/jmoiron/sqlx"
	"github.com/super667/user/common/casbinx"
	"github.com/super667/user/common/ldap"
	"github.com/super667/user/model"
	"github.com/super667/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	RoleModel     model.RoleModel
	PermModel     model.PermissionModel
	StrategyModel model.StrategyModel
	UserRoleModel model.UserRoleModel
	TokenModel    model.TokenModel
	LdapPool      *ldap.Pool
	Rds           *redis.Redis
	Casbin        *casbinx.Casbin
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySql.DataSource)
	// connect to the database first.
	db, err := casbinsqlx.Connect("mysql", c.MySql.DataSource)
	if err != nil {
		panic(err)
	}

	casbin := casbinx.New(db)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(conn),
		RoleModel:     model.NewRoleModel(conn),
		PermModel:     model.NewPermissionModel(conn),
		StrategyModel: model.NewStrategyModel(conn),
		UserRoleModel: model.NewUserRoleModel(conn),
		TokenModel:    model.NewTokenModel(conn, c.CacheRedis, cache.WithExpiry(time.Duration(5))),
		LdapPool:      ldap.NewLdapPool(c.Ldap),
		Rds:           redis.MustNewRedis(c.CacheRedis[0].RedisConf),
		Casbin:        casbin,
	}
}
