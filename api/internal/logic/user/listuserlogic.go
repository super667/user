package user

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser(req *types.ListUserReq) (resp *types.ListUserResp, err error) {
	res, err := l.svcCtx.UserRpc.ListUser(l.ctx, &userclient.ListUserReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return
	}
	userInfos := make([]types.UserDetail, 0)
	copier.Copy(&userInfos, res.UserDetail)
	resp = &types.ListUserResp{
		Users: userInfos,
		ListResp: types.ListResp{
			Page:     res.Page,
			PageSize: res.PageSize,
			Total:    res.Total,
		},
	}
	return
}
