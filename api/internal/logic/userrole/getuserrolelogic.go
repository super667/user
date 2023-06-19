package userrole

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/userroleservice"

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
	res, err := l.svcCtx.UserRoleRpc.GetUserRoleById(l.ctx, &userroleservice.GetUserRoleByIdReq{
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
