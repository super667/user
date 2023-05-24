package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPermLogic {
	return &ListPermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPermLogic) ListPerm(in *user.ListPermReq) (*user.ListPermResp, error) {
	// todo: add your logic here and delete this line

	return &user.ListPermResp{}, nil
}
