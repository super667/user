package auth

import (
	"context"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WeComLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWeComLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WeComLoginLogic {
	return &WeComLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WeComLoginLogic) WeComLogin(req *types.WeComLoginReq) (resp *types.WeComLoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
