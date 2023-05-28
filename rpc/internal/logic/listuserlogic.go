package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserLogic) ListUser(in *user.ListUserReq) (*user.ListUserResp, error) {
	resp := &user.ListUserResp{}
	res, err := l.svcCtx.UserModel.FindManyByPage(l.ctx, in.Search, in.Page, in.PageSize)
	if err != nil {
		return resp, err
	}
	userDetails := make([]*user.UserDetail, 0)
	err = copier.Copy(&userDetails, res)
	if err != nil {
		return resp, err
	}

	total, err := l.svcCtx.UserModel.Count(l.ctx, in.Search)
	if err != nil {
		return resp, err
	}
	return &user.ListUserResp{
		Page:       in.Page,
		PageSize:   in.PageSize,
		Total:      total,
		UserDetail: userDetails,
	}, nil
}
