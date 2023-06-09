package roleservicelogic

import (
	"context"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *user.DeleteRoleReq) (*user.DeleteRoleResp, error) {
	err := l.svcCtx.RoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return &user.DeleteRoleResp{}, err
	}
	return &user.DeleteRoleResp{}, nil
}
