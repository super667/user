package role

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.UpdateRoleResp, err error) {
	var updateInfo = &userclient.RoleInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.UpdateRole(l.ctx, &userclient.UpdateRoleReq{
		Id:       req.Id,
		RoleInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
