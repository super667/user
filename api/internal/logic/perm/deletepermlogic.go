package perm

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermLogic {
	return &DeletePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePermLogic) DeletePerm(req *types.DeletePermReq) (resp *types.DeletePermResp, err error) {
	// todo: add your logic here and delete this line

	return
}
