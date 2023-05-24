package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

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
	// todo: add your logic here and delete this line

	return &user.ListStrategyResp{}, nil
}
