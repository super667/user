package role

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

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
	var updateInfo = &userclient.RoleInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.PatchRole(l.ctx, &userclient.PatchRoleReq{
		Id:       req.Id,
		RoleInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
