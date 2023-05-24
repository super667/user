package userrole

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserRoleLogic {
	return &ListUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserRoleLogic) ListUserRole(req *types.ListUserRoleReq) (resp *types.ListUserRoleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
