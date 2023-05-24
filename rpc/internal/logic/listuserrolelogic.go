package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserRoleLogic {
	return &ListUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserRoleLogic) ListUserRole(in *user.ListUserRoleReq) (*user.ListUserRoleResp, error) {
	// todo: add your logic here and delete this line

	return &user.ListUserRoleResp{}, nil
}
