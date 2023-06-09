package userservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchUserLogic {
	return &PatchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchUserLogic) PatchUser(in *user.PatchUserReq) (*user.PatchUserResp, error) {
	var userInfo model.User
	copier.Copy(&userInfo, in.UserInfo)
	userInfo.Id = in.Id
	err := l.svcCtx.UserModel.PartialUpdate(l.ctx, &userInfo)
	if err != nil {
		return &user.PatchUserResp{}, err
	}

	return &user.PatchUserResp{}, nil
}
