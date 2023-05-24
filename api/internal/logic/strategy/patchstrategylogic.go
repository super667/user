package strategy

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchStrategyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchStrategyLogic {
	return &PatchStrategyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchStrategyLogic) PatchStrategy(req *types.PatchStrategyReq) (resp *types.PatchStrategyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
