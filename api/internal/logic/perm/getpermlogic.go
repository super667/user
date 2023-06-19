package perm

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/permservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermLogic {
	return &GetPermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermLogic) GetPerm(req *types.GetPermReq) (resp *types.GetPermResp, err error) {
	res, err := l.svcCtx.PermRpc.GetPermById(l.ctx, &permservice.GetPermByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return resp, err
	}
	var permDetail types.PermDetail
	copier.Copy(&permDetail, res.PermDetail)
	resp = &types.GetPermResp{PermDetail: permDetail}
	return
}
