package userrole

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(req *types.UpdateUserRoleReq) (resp *types.UpdateUserRoleResp, err error) {
	var updateInfo = &userclient.UserRoleInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.UpdateUserRole(l.ctx, &userclient.UpdateUserRoleReq{
		Id:           req.Id,
		UserRoleInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
