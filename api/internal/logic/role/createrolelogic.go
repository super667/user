package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/rpc/userclient"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleReq) (resp *types.CreateRoleResp, err error) {
	var u = &userclient.CreateRoleReq{
		RoleInfo: &userclient.RoleInfo{},
	}
	copier.Copy(u.RoleInfo, req)
	res, err := l.svcCtx.UserRpc.CreateRole(l.ctx, u)
	if err != nil {
		return
	}
	resp = &types.CreateRoleResp{Id: res.Id}

	return
}
