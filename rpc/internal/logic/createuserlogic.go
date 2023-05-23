package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"user/model"

	"user/rpc/internal/svc"
	"user/rpc/user"

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
