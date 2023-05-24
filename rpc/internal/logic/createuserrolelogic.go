package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

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
	// todo: add your logic here and delete this line

	return &user.CreateUserRoleResp{}, nil
}
