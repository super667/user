package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"user/model"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *user.UpdateRoleReq) (*user.UpdateRoleResp, error) {
	var roleInfo model.Role
	copier.Copy(&roleInfo, in.RoleInfo)
	roleInfo.Id = in.Id
	err := l.svcCtx.RoleModel.Update(l.ctx, &roleInfo)
	if err != nil {
		return &user.UpdateRoleResp{}, err
	}

	return &user.UpdateRoleResp{}, nil
}