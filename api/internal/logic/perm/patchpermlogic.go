package perm

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchPermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchPermLogic {
	return &PatchPermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchPermLogic) PatchPerm(req *types.PatchPermReq) (resp *types.PatchPermResp, err error) {
	// todo: add your logic here and delete this line

	return
}
