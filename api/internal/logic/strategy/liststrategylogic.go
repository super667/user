package strategy

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStrategyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStrategyLogic {
	return &ListStrategyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListStrategyLogic) ListStrategy(req *types.ListStrategyReq) (resp *types.ListStrategyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
