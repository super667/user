package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/rpc/userclient"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	var u = &userclient.CreateUserReq{
		UserInfo: &userclient.UserInfo{},
	}
	copier.Copy(u.UserInfo, req)
	res, err := l.svcCtx.UserRpc.CreateUser(l.ctx, u)
	if err != nil {
		return
	}
	resp = &types.CreateUserResp{Id: res.Id}

	return
}
