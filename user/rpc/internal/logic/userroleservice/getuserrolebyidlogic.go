package userroleservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/common/errorx"
	"github.com/super667/user/user/model"

	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoleByIdLogic {
	return &GetUserRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRoleByIdLogic) GetUserRoleById(in *user.GetUserRoleByIdReq) (*user.GetUserRoleByIdResp, error) {
	Info, err := l.svcCtx.UserRoleModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("不存在")
	default:
		return nil, err
	}
	var userRoleDetail user.UserRoleDetail
	copier.Copy(&userRoleDetail, Info)

	return &user.GetUserRoleByIdResp{UserRoleDetail: &userRoleDetail}, nil
}
