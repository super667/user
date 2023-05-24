package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &user.UpdateUserRoleResp{}, nil
}
