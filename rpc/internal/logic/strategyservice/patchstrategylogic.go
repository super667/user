package strategyservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

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
	var strategyInfo model.Strategy
	copier.Copy(&strategyInfo, in.StrategyInfo)
	strategyInfo.Id = in.Id
	err := l.svcCtx.StrategyModel.PartialUpdate(l.ctx, &strategyInfo)
	if err != nil {
		return &user.PatchStrategyResp{}, err
	}

	return &user.PatchStrategyResp{}, nil
}
