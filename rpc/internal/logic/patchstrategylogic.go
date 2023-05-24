package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchStrategyLogic {
	return &PatchStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchStrategyLogic) PatchStrategy(in *user.PatchStrategyReq) (*user.PatchStrategyResp, error) {
	// todo: add your logic here and delete this line

	return &user.PatchStrategyResp{}, nil
}
