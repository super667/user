package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"user/model"

	"user/rpc/internal/svc"
	"user/rpc/user"

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
