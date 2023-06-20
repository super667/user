package roleservicelogic

import (
	"context"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveRoleForUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveRoleForUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRoleForUserLogic {
	return &RemoveRoleForUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveRoleForUserLogic) RemoveRoleForUser(in *user.RemoveRoleForUserReq) (*user.RemoveRoleForUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.RemoveRoleForUserResp{}, nil
}
