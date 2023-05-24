package perm

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPermLogic {
	return &ListPermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPermLogic) ListPerm(req *types.ListPermReq) (resp *types.ListPermResp, err error) {
	// todo: add your logic here and delete this line

	return
}
