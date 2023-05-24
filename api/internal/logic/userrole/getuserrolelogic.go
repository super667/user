package userrole

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoleLogic {
	return &GetUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRoleLogic) GetUserRole(req *types.GetUserRoleReq) (resp *types.GetUserRoleResp, err error) {
	res, err := l.svcCtx.UserRpc.GetUserRoleById(l.ctx, &userclient.GetUserRoleByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return resp, err
	}
	var userRoleDetail types.UserRoleDetail
	copier.Copy(&userRoleDetail, res.UserRoleDetail)
	resp = &types.GetUserRoleResp{UserRoleDetail: userRoleDetail}
	return
}
