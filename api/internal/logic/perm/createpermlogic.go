package perm

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/rpc/userclient"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermLogic {
	return &CreatePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePermLogic) CreatePerm(req *types.CreatePermReq) (resp *types.CreatePermResp, err error) {
	var u = &userclient.CreatePermReq{
		PermInfo: &userclient.PermInfo{},
	}
	copier.Copy(u.PermInfo, req)
	res, err := l.svcCtx.UserRpc.CreatePerm(l.ctx, u)
	if err != nil {
		return
	}
	resp = &types.CreatePermResp{Id: res.Id}

	return
}
