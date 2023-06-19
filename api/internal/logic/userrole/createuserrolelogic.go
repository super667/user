package userrole

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/userroleservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserRoleLogic {
	return &CreateUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserRoleLogic) CreateUserRole(req *types.CreateUserRoleReq) (resp *types.CreateUserRoleResp, err error) {
	var u = &userroleservice.CreateUserRoleReq{
		UserRoleInfo: &userroleservice.UserRoleInfo{},
	}
	copier.Copy(u.UserRoleInfo, req)
	res, err := l.svcCtx.UserRoleRpc.CreateUserRole(l.ctx, u)
	if err != nil {
		return
	}
	resp = &types.CreateUserRoleResp{Id: res.Id}

	return
}
