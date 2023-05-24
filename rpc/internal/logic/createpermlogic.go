package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermLogic {
	return &CreatePermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePermLogic) CreatePerm(in *user.CreatePermReq) (*user.CreatePermResp, error) {
	// todo: add your logic here and delete this line

	return &user.CreatePermResp{}, nil
}
