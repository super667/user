package userrole

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(req *types.UpdateUserRoleReq) (resp *types.UpdateUserRoleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
