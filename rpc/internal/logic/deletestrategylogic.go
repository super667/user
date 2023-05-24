package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

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
	// todo: add your logic here and delete this line

	return &user.DeleteStrategyResp{}, nil
}
