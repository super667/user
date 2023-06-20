package roleservicelogic

import (
	"context"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRoleUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleUserLogic {
	return &ListRoleUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListRoleUserLogic) ListRoleUser(in *user.ListRoleForUserReq) (*user.ListRoleForUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.ListRoleForUserResp{}, nil
}
