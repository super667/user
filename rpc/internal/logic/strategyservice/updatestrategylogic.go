package strategyservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStrategyLogic {
	return &UpdateStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStrategyLogic) UpdateStrategy(in *user.UpdateStrategyReq) (*user.UpdateStrategyResp, error) {
	var strategyInfo model.Strategy
	copier.Copy(&strategyInfo, in.StrategyInfo)
	strategyInfo.Id = in.Id
	err := l.svcCtx.StrategyModel.Update(l.ctx, &strategyInfo)
	if err != nil {
		return &user.UpdateStrategyResp{}, err
	}

	return &user.UpdateStrategyResp{}, nil
}
