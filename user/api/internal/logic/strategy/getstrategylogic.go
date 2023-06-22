package strategy

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/strategyservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStrategyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStrategyLogic {
	return &GetStrategyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStrategyLogic) GetStrategy(req *types.GetStrategyReq) (resp *types.GetStrategyResp, err error) {
	res, err := l.svcCtx.StrategyRpc.GetStrategyById(l.ctx, &strategyservice.GetStrategyByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return resp, err
	}
	var strategyDetail types.StrategyDetail
	copier.Copy(&strategyDetail, res.StrategyDetail)
	resp = &types.GetStrategyResp{StrategyDetail: strategyDetail}
	return
}
