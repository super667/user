package strategy

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStrategyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStrategyLogic {
	return &CreateStrategyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStrategyLogic) CreateStrategy(req *types.CreateStrategyReq) (resp *types.CreateStrategyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
