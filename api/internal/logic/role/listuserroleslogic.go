package role

import (
	"context"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserRolesLogic {
	return &ListUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserRolesLogic) ListUserRoles(req *types.ListUserRolesReq) (resp *types.ListUserRolesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
