package authservicelogic

import (
	"context"
	"github.com/super667/user/common/ctxdata"
	"github.com/super667/user/user/model"
	"github.com/super667/user/user/rpc/client/userservice"

	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.LogoutResp, error) {
	name := ctxdata.GetNameFromCtx(l.ctx)
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, name)
	if err != nil {
		return &userservice.LogoutResp{}, err
	}
	_, err = l.svcCtx.TokenModel.Insert(l.ctx, &model.Token{
		UserId: userInfo.Id,
		Token:  in.Token,
	})
	if err != nil {
		return &userservice.LogoutResp{}, err
	}

	return &userservice.LogoutResp{}, nil
}
