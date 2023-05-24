package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchPermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchPermLogic {
	return &PatchPermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchPermLogic) PatchPerm(in *user.PatchPermReq) (*user.PatchPermResp, error) {
	// todo: add your logic here and delete this line

	return &user.PatchPermResp{}, nil
}
