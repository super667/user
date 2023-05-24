package userrole

import (
	"context"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserRoleLogic {
	return &DeleteUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserRoleLogic) DeleteUserRole(req *types.DeleteUserRoleReq) (resp *types.DeleteUserRoleResp, err error) {
	_, err = l.svcCtx.UserRpc.DeleteUserRole(l.ctx, &userclient.DeleteUserRoleReq{
		Id: req.Id,
	})
	if err != nil {
		return
	}

	return
}