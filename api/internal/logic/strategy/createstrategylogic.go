package strategy

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/strategyservice"

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
	var u = &strategyservice.CreateStrategyReq{
		StrategyInfo: &strategyservice.StrategyInfo{},
	}
	copier.Copy(u.StrategyInfo, req)
	res, err := l.svcCtx.StrategyRpc.CreateStrategy(l.ctx, u)
	if err != nil {
		return
	}
	resp = &types.CreateStrategyResp{Id: res.Id}

	return
}
