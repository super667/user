package userrole

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/userroleservice"

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
	var updateInfo = &userroleservice.UserRoleInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRoleRpc.UpdateUserRole(l.ctx, &userroleservice.UpdateUserRoleReq{
		Id:           req.Id,
		UserRoleInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
