package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/common/errorx"
	"github.com/super667/user/model"
	"strings"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNumberLogic {
	return &GetUserByNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNumberLogic) GetUserByNumber(in *user.GetUserByNumberReq) (*user.GetUserByNumberResp, error) {
	if len(strings.TrimSpace(in.Number)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}
	userInfoDB, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, in.Number)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户不存在")
	default:
		return nil, err
	}
	var userDetail user.UserDetail
	copier.Copy(&userDetail, userInfoDB)
	return &user.GetUserByNumberResp{UserDetail: &userDetail}, nil
}
