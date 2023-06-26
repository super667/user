package authservicelogic

import (
	"context"
	"github.com/super667/user/user/model"
	"github.com/super667/user/user/rpc/client/authservice"

	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlackListTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlackListTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlackListTokenLogic {
	return &BlackListTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlackListTokenLogic) BlackListToken(in *user.BlackListTokenReq) (*user.BlackListTokenResp, error) {
	resp := &authservice.BlackListTokenResp{}
	res, err := l.svcCtx.TokenModel.FindOneByToken(l.ctx, in.Token)
	if err == model.ErrNotFound {
		resp.Invalid = false
		return resp, nil
	}
	if res != nil {
		resp.Invalid = true
		return resp, nil
	}
	resp.Invalid = false
	return resp, nil
}
