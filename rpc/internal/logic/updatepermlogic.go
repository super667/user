package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermLogic {
	return &UpdatePermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePermLogic) UpdatePerm(in *user.UpdatePermReq) (*user.UpdatePermResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdatePermResp{}, nil
}
