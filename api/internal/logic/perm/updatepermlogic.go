package perm

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermLogic {
	return &UpdatePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePermLogic) UpdatePerm(req *types.UpdatePermReq) (resp *types.UpdatePermResp, err error) {
	// todo: add your logic here and delete this line

	return
}
