package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchRoleLogic {
	return &PatchRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchRoleLogic) PatchRole(in *user.PatchRoleReq) (*user.PatchRoleResp, error) {
	// todo: add your logic here and delete this line

	return &user.PatchRoleResp{}, nil
}
