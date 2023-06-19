package svc

import (
	"github.com/super667/user/api/internal/config"
	"github.com/super667/user/api/internal/middleware"
	"github.com/super667/user/model"
	"github.com/super667/user/rpc/client/authservice"
	"github.com/super667/user/rpc/client/permservice"
	"github.com/super667/user/rpc/client/roleservice"
	"github.com/super667/user/rpc/client/strategyservice"
	"github.com/super667/user/rpc/client/userroleservice"
	"github.com/super667/user/rpc/client/userservice"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userservice.UserService
	AuthRpc     authservice.AuthService
	PermRpc     permservice.PermService
	RoleRpc     roleservice.RoleService
	StrategyRpc strategyservice.StrategyService
	UserRoleRpc userroleservice.UserRoleService
	Email       rest.Middleware
	TokenModel  model.TokenModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySql.DataSource)
	return &ServiceContext{
		Config:      c,
		UserRpc:     userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		AuthRpc:     authservice.NewAuthService(zrpc.MustNewClient(c.UserRpc)),
		PermRpc:     permservice.NewPermService(zrpc.MustNewClient(c.UserRpc)),
		RoleRpc:     roleservice.NewRoleService(zrpc.MustNewClient(c.UserRpc)),
		StrategyRpc: strategyservice.NewStrategyService(zrpc.MustNewClient(c.UserRpc)),
		UserRoleRpc: userroleservice.NewUserRoleService(zrpc.MustNewClient(c.UserRpc)),
		Email:       middleware.EmailWithUser(model.NewUserModel(conn)),
	}
}
