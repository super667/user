package userrole

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserRoleLogic {
	return &CreateUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserRoleLogic) CreateUserRole(req *types.CreateUserRoleReq) (resp *types.CreateUserRoleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
