package auth

import (
	"context"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.AuthRpc.Login(l.ctx, &userservice.LoginReq{
		Account:  req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  res.AccessToken,
		AccessExpire: res.AccessExpire,
		RefreshToken: res.RefreshToken,
	}, nil
}
