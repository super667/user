package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/model"
	"user/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	RoleModel     model.RoleModel
	PermModel     model.PermissionModel
	StrategyModel model.StrategyModel
	UserRoleModel model.UserRoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(conn),
		RoleModel:     model.NewRoleModel(conn),
		PermModel:     model.NewPermissionModel(conn),
		StrategyModel: model.NewStrategyModel(conn),
		UserRoleModel: model.NewUserRoleModel(conn),
	}
}
