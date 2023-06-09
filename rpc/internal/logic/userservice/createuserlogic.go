package userservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/common/cryptx"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	resp := &user.CreateUserResp{}
	userInfo := &model.User{}
	err := copier.Copy(userInfo, in.UserInfo)
	if err != nil {
		l.Logger.Error(err)
		return &user.CreateUserResp{}, err
	}
	userInfo.Password = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, "123456")
	insertRes, err := l.svcCtx.UserModel.Insert(l.ctx, userInfo)
	if err != nil {
		return resp, err
	}
	lastId, err := insertRes.LastInsertId()
	if err != nil {
		return resp, err
	}

	resp.Id = lastId
	return resp, nil
}
