package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserRoleLogic {
	return &ListUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserRoleLogic) ListUserRole(in *user.ListUserRoleReq) (*user.ListUserRoleResp, error) {
	resp := &user.ListUserRoleResp{}
	res, err := l.svcCtx.PermModel.FindManyByPage(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return resp, err
	}
	userRoleDetails := make([]*user.UserRoleDetail, 0)
	err = copier.Copy(&userRoleDetails, res)
	if err != nil {
		return resp, err
	}

	total, err := l.svcCtx.StrategyModel.Count(l.ctx)
	if err != nil {
		return resp, err
	}

	return &user.ListUserRoleResp{
		Page:           in.Page,
		PageSize:       in.PageSize,
		Total:          total,
		UserRoleDetail: userRoleDetails,
	}, nil
}
