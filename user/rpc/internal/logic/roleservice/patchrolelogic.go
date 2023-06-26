package roleservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/model"

	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchRoleLogic {
	return &PatchRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchRoleLogic) PatchRole(in *user.PatchRoleReq) (*user.PatchRoleResp, error) {
	var roleInfo model.Role
	copier.Copy(&roleInfo, in.RoleInfo)
	roleInfo.Id = in.Id
	err := l.svcCtx.RoleModel.PartialUpdate(l.ctx, &roleInfo)
	if err != nil {
		return &user.PatchRoleResp{}, err
	}
	return &user.PatchRoleResp{}, nil
}
