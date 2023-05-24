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
	err := l.svcCtx.UserRoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return &user.DeleteUserRoleResp{}, err
	}
	return &user.DeleteUserRoleResp{}, nil
}
