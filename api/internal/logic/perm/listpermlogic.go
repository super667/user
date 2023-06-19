package perm

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/permservice"

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
	res, err := l.svcCtx.PermRpc.ListPerm(l.ctx, &permservice.ListPermReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return
	}
	permDetails := make([]types.PermDetail, 0)
	copier.Copy(&permDetails, res.PermDetail)
	resp = &types.ListPermResp{
		Perms: permDetails,
		ListResp: types.ListResp{
			Page:     res.Page,
			PageSize: res.PageSize,
			Total:    res.Total,
		},
	}

	return
}
