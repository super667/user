package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

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
	// todo: add your logic here and delete this line

	return &user.UpdateStrategyResp{}, nil
}
