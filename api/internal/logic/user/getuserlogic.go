package user

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	res, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &userclient.GetUserByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return resp, err
	}
	var userDetail types.UserDetail
	copier.Copy(&userDetail, res.UserDetail)
	resp = &types.GetUserResp{UserDetail: userDetail}
	return
}
