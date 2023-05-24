package auth

import (
	"context"
	"github.com/super667/user/rpc/userclient"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

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
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterReq{
		Name:       req.Name,
		GenderCode: req.GenderCode,
		Phone:      req.Phone,
		Password:   req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Id:         res.Id,
		Name:       res.Name,
		GenderCode: res.GenderCode,
		Phone:      res.Phone,
	}, nil
}
