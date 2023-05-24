package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermLogic {
	return &DeletePermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePermLogic) DeletePerm(in *user.DeletePermReq) (*user.DeletePermResp, error) {
	err := l.svcCtx.PermModel.Delete(l.ctx, in.Id)
	if err != nil {
		return &user.DeletePermResp{}, err
	}
	return &user.DeletePermResp{}, nil
}
