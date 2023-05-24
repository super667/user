package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"user/model"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(in *user.UpdateUserRoleReq) (*user.UpdateUserRoleResp, error) {
	var userRoleInfo model.UserRole
	copier.Copy(&userRoleInfo, in.UserRoleInfo)
	userRoleInfo.Id = in.Id
	err := l.svcCtx.UserRoleModel.Update(l.ctx, &userRoleInfo)
	if err != nil {
		return &user.UpdateUserRoleResp{}, err
	}

	return &user.UpdateUserRoleResp{}, nil
}
