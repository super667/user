package user

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/userservice"

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
	res, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &userservice.GetUserByIdReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Error("GetUser", err.Error())
		return resp, err
	}
	var userDetail types.UserDetail
	copier.Copy(&userDetail, res.UserDetail)
	resp = &types.GetUserResp{UserDetail: userDetail}
	fmt.Printf("获取到user:%#v", resp)
	return
}
