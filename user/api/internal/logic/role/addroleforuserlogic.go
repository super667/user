package role

import (
	"context"

	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleForUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleForUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleForUserLogic {
	return &AddRoleForUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleForUserLogic) AddRoleForUser(req *types.AddRoleForUserReq) (resp *types.AddRoleForUserResp, err error) {

	return
}
