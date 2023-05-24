package userrole

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

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
	var u = &userclient.CreateUserRoleReq{
		UserRoleInfo: &userclient.UserRoleInfo{},
	}
	copier.Copy(u.UserRoleInfo, req)
	res, err := l.svcCtx.UserRpc.CreateUserRole(l.ctx, u)
	if err != nil {
		return
	}
	resp = &types.CreateUserRoleResp{Id: res.Id}

	return
}
