package userrole

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/rpc/userclient"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserRoleLogic {
	return &ListUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserRoleLogic) ListUserRole(req *types.ListUserRoleReq) (resp *types.ListUserRoleResp, err error) {
	res, err := l.svcCtx.UserRpc.ListUserRole(l.ctx, &userclient.ListUserRoleReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return
	}
	userRoleDetails := make([]types.UserRoleDetail, 0)
	copier.Copy(&userRoleDetails, res.UserRoleDetail)
	resp = &types.ListUserRoleResp{
		UserRoles: userRoleDetails,
		ListResp: types.ListResp{
			Page:     res.Page,
			PageSize: res.PageSize,
			Total:    res.Total,
		},
	}
	return
}
