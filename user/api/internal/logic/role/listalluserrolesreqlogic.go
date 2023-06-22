package role

import (
	"context"

	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllUserRolesReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAllUserRolesReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllUserRolesReqLogic {
	return &ListAllUserRolesReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAllUserRolesReqLogic) ListAllUserRolesReq(req *types.ListAllUserRolesReq) (resp *types.ListAllUserRolesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
