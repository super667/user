package role

import (
	"context"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleForUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleForUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleForUserLogic {
	return &UpdateRoleForUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleForUserLogic) UpdateRoleForUser(req *types.RemoveRoleForUserReq) (resp *types.RemoveRoleForUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
