package auth

import (
	"context"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/authservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	res, err := l.svcCtx.AuthRpc.Register(l.ctx, &authservice.RegisterReq{
		UserName: req.UserName,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Id:       res.Id,
		UserName: res.UserName,
		Phone:    res.Phone,
	}, nil
}
