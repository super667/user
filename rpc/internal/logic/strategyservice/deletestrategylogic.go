package strategyservicelogic

import (
	"context"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStrategyLogic {
	return &DeleteStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStrategyLogic) DeleteStrategy(in *user.DeleteStrategyReq) (*user.DeleteStrategyResp, error) {
	err := l.svcCtx.StrategyModel.Delete(l.ctx, in.Id)
	if err != nil {
		return &user.DeleteStrategyResp{}, err
	}
	return &user.DeleteStrategyResp{}, nil
}
