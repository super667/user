package strategy

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStrategyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStrategyLogic {
	return &DeleteStrategyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStrategyLogic) DeleteStrategy(req *types.DeleteStrategyReq) (resp *types.DeleteStrategyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
