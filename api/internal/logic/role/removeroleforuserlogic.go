package role

import (
	"context"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveRoleForUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveRoleForUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRoleForUserLogic {
	return &RemoveRoleForUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveRoleForUserLogic) RemoveRoleForUser(req *types.RemoveRoleForUserReq) (resp *types.RemoveRoleForUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
