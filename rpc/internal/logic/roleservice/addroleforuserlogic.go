package roleservicelogic

import (
	"context"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleForUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleForUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleForUserLogic {
	return &AddRoleForUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRoleForUserLogic) AddRoleForUser(in *user.AddRoleForUserReq) (*user.AddRoleForUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddRoleForUserResp{}, nil
}
