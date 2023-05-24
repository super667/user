package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserRoleLogic {
	return &DeleteUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserRoleLogic) DeleteUserRole(in *user.DeleteUserRoleReq) (*user.DeleteUserRoleResp, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteUserRoleResp{}, nil
}
