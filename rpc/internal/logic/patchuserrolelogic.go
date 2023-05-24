package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchUserRoleLogic {
	return &PatchUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchUserRoleLogic) PatchUserRole(in *user.PatchUserRoleReq) (*user.PatchUserRoleResp, error) {
	// todo: add your logic here and delete this line

	return &user.PatchUserRoleResp{}, nil
}
