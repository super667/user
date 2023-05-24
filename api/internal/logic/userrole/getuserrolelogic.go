package userrole

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoleLogic {
	return &GetUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRoleLogic) GetUserRole(req *types.GetUserRoleReq) (resp *types.GetUserRoleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
