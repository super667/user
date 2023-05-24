package perm

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

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
	res, err := l.svcCtx.UserRpc.GetPermById(l.ctx, &userclient.GetPermByIdReq{
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
