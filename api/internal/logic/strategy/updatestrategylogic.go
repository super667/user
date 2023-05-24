package strategy

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStrategyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStrategyLogic {
	return &UpdateStrategyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStrategyLogic) UpdateStrategy(req *types.UpdateStrategyReq) (resp *types.UpdateStrategyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
