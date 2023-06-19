package auth

import (
	"context"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/authservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.RefreshTokenResp, err error) {
	resp = &types.RefreshTokenResp{}
	res, err := l.svcCtx.AuthRpc.FreshToken(l.ctx, &authservice.RefreshTokenReq{})
	if err != nil {
		return
	}
	resp.AccessToken = res.AccessToken
	resp.AccessExpire = res.AccessExpire
	return
}
