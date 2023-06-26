package strategy

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/strategyservice"

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
	res, err := l.svcCtx.StrategyRpc.ListStrategy(l.ctx, &strategyservice.ListStrategyReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return
	}
	strategyDetails := make([]types.StrategyDetail, 0)
	copier.Copy(&strategyDetails, res.StrategyDetail)
	resp = &types.ListStrategyResp{
		Strategys: strategyDetails,
		ListResp: types.ListResp{
			Page:     res.Page,
			PageSize: res.PageSize,
			Total:    res.Total,
		},
	}
	return
}
