package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/roleservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchRoleLogic {
	return &PatchRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchRoleLogic) PatchRole(req *types.PatchRoleReq) (resp *types.PatchRoleResp, err error) {
	var updateInfo = &roleservice.RoleInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.RoleRpc.PatchRole(l.ctx, &roleservice.PatchRoleReq{
		Id:       req.Id,
		RoleInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
