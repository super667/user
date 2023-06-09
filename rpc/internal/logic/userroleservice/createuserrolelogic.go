package userroleservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserRoleLogic {
	return &CreateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserRoleLogic) CreateUserRole(in *user.CreateUserRoleReq) (*user.CreateUserRoleResp, error) {
	resp := &user.CreateUserRoleResp{}
	userRoleInfo := &model.UserRole{}
	err := copier.Copy(userRoleInfo, in.UserRoleInfo)
	if err != nil {
		l.Logger.Error(err)
		return resp, err
	}
	insertRes, err := l.svcCtx.UserRoleModel.Insert(l.ctx, userRoleInfo)
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
