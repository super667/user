package strategyservicelogic

import (
	"context"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthenticateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateLogic {
	return &AuthenticateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthenticateLogic) Authenticate(in *user.AuthenticateReq) (*user.AuthenticateResp, error) {
	// todo: add your logic here and delete this line

	return &user.AuthenticateResp{}, nil
}
