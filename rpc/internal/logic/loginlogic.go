package logic

import (
	"context"
	"database/sql"
	"github.com/super667/user/common/cryptx"
	"github.com/super667/user/common/jwtx"
	"github.com/super667/user/model"
	"google.golang.org/grpc/status"
	"time"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

const FreshTokenExpires = 60 * 60 * 24

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	resp := &user.LoginResp{}
	res, err := l.svcCtx.UserModel.GetOne(l.ctx, in.Account)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return resp, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	now := time.Now()
	nowTimeStamp := now.Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	refreshExpire := l.svcCtx.Config.JWT.RefreshExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, nowTimeStamp, accessExpire, res.UserName)
	if err != nil {
		return resp, err
	}
	refreshToken, err := jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, nowTimeStamp, refreshExpire, res.UserName)
	if err != nil {
		return resp, err
	}

	_, err = l.svcCtx.TokenModel.Insert(l.ctx, &model.Token{
		UserId: res.Id,
		Token:  refreshToken,
		ExpiredAt: sql.NullTime{
			Time:  now.Add(time.Duration(FreshTokenExpires)),
			Valid: true,
		},
	})
	if err != nil {
		return resp, err
	}

	return &user.LoginResp{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
		RefreshToken: refreshToken,
	}, nil
}
