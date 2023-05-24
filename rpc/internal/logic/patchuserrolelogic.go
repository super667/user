package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchUserRoleLogic {
	return &PatchUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchUserRoleLogic) PatchUserRole(in *user.PatchUserRoleReq) (*user.PatchUserRoleResp, error) {
	var userRoleInfo model.UserRole
	copier.Copy(&userRoleInfo, in.UserRoleInfo)
	userRoleInfo.Id = in.Id
	err := l.svcCtx.UserRoleModel.PartialUpdate(l.ctx, &userRoleInfo)
	if err != nil {
		return &user.PatchUserRoleResp{}, err
	}

	return &user.PatchUserRoleResp{}, nil
}
