package authservicelogic

import (
	"context"
	"github.com/super667/user/common/ctxdata"
	"github.com/super667/user/common/jwtx"
	"time"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreshTokenLogic {
	return &FreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FreshTokenLogic) FreshToken(in *user.RefreshTokenReq) (*user.RefreshTokenResp, error) {
	now := time.Now().Unix()
	name := ctxdata.GetNameFromCtx(l.ctx)
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, name)
	if err != nil {
		return nil, err
	}

	return &user.RefreshTokenResp{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
	}, nil
}
