package roleservicelogic

import (
	"context"

	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleForUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleForUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleForUserLogic {
	return &UpdateRoleForUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleForUserLogic) UpdateRoleForUser(in *user.UpdateRoleForUserReq) (*user.UpdateRoleForUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateRoleForUserResp{}, nil
}
