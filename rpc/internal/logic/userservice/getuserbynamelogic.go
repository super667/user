package userservicelogic

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

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *user.GetUserByNameReq) (*user.GetUserByNameResp, error) {
	if len(strings.TrimSpace(in.Name)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}
	userInfoDB, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, in.Name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}
	var userDetail user.UserDetail
	copier.Copy(&userDetail, userInfoDB)
	return &user.GetUserByNameResp{UserDetail: &userDetail}, nil
}
