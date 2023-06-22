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

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser(req *types.ListUserReq) (resp *types.ListUserResp, err error) {
	l.Logger.Info(fmt.Sprintf("ListUser: 获取用户列表：%+v", req.ListReq))
	resp = &types.ListUserResp{
		ListResp: types.ListResp{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}
	res, err := l.svcCtx.UserRpc.ListUser(l.ctx, &userservice.ListUserReq{
		Search:   req.Search,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return
	}
	userInfos := make([]types.UserDetail, 0)
	copier.Copy(&userInfos, res.UserDetail)
	resp.Users = userInfos
	resp.Total = res.Total
	return
}
