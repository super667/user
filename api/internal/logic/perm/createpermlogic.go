package perm

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermLogic {
	return &CreatePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePermLogic) CreatePerm(req *types.CreatePermReq) (resp *types.CreatePermResp, err error) {
	// todo: add your logic here and delete this line

	return
}
