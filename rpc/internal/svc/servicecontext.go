package svc

import (
	"github.com/super667/user/common/ldap"
	"github.com/super667/user/model"
	"github.com/super667/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	RoleModel     model.RoleModel
	PermModel     model.PermissionModel
	StrategyModel model.StrategyModel
	UserRoleModel model.UserRoleModel
	LdapPool      *ldap.Pool
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
		LdapPool:      ldap.NewLdapPool(c.Ldap),
	}
}
