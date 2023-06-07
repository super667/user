package svc

import (
	"github.com/super667/user/api/internal/config"
	"github.com/super667/user/api/internal/middleware"
	"github.com/super667/user/model"
	"github.com/super667/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc    userclient.User
	Email      rest.Middleware
	TokenModel model.TokenModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySql.DataSource)
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Email:   middleware.EmailWithUser(model.NewUserModel(conn)),
	}
}
