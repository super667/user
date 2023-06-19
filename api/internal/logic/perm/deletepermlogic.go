package perm

import (
	"context"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/permservice"

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
	_, err = l.svcCtx.PermRpc.DeletePerm(l.ctx, &permservice.DeletePermReq{
		Id: req.Id,
	})
	if err != nil {
		return
	}

	return
}
