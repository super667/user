package roleservicelogic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListRoleLogic) ListRole(in *user.ListRoleReq) (*user.ListRoleResp, error) {
	resp := &user.ListRoleResp{}
	res, err := l.svcCtx.PermModel.FindManyByPage(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return resp, err
	}
	roleDetails := make([]*user.RoleDetail, 0)
	err = copier.Copy(&roleDetails, res)
	if err != nil {
		return resp, err
	}

	total, err := l.svcCtx.RoleModel.Count(l.ctx)
	if err != nil {
		return resp, err
	}

	return &user.ListRoleResp{
		Page:       in.Page,
		PageSize:   in.PageSize,
		Total:      total,
		RoleDetail: roleDetails,
	}, nil
}
