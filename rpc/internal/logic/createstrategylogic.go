package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStrategyLogic {
	return &CreateStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStrategyLogic) CreateStrategy(in *user.CreateStrategyReq) (*user.CreateStrategyResp, error) {
	// todo: add your logic here and delete this line

	return &user.CreateStrategyResp{}, nil
}
