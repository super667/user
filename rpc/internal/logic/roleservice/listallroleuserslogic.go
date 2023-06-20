package roleservicelogic

import (
	"context"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllRoleUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAllRoleUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllRoleUsersLogic {
	return &ListAllRoleUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListAllRoleUsersLogic) ListAllRoleUsers(in *user.ListAllRoleUsersReq) (*user.ListAllRoleUsersResp, error) {
	// todo: add your logic here and delete this line

	return &user.ListAllRoleUsersResp{}, nil
}
