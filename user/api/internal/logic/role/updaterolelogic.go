package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/roleservice"

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
	var updateInfo = &roleservice.RoleInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.RoleRpc.UpdateRole(l.ctx, &roleservice.UpdateRoleReq{
		Id:       req.Id,
		RoleInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
