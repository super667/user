package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/model"
	"user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
	}
}
