package strategyservicelogic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStrategyLogic {
	return &ListStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListStrategyLogic) ListStrategy(in *user.ListStrategyReq) (*user.ListStrategyResp, error) {
	resp := &user.ListStrategyResp{}
	res, err := l.svcCtx.PermModel.FindManyByPage(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return resp, err
	}
	strategyDetails := make([]*user.StrategyDetail, 0)
	err = copier.Copy(&strategyDetails, res)
	if err != nil {
		return resp, err
	}

	total, err := l.svcCtx.StrategyModel.Count(l.ctx)
	if err != nil {
		return resp, err
	}

	return &user.ListStrategyResp{
		Page:           in.Page,
		PageSize:       in.PageSize,
		Total:          total,
		StrategyDetail: strategyDetails,
	}, nil

	return &user.ListStrategyResp{}, nil
}
