package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPermLogic {
	return &ListPermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPermLogic) ListPerm(in *user.ListPermReq) (*user.ListPermResp, error) {
	resp := &user.ListPermResp{}
	res, err := l.svcCtx.PermModel.FindManyByPage(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return resp, err
	}
	permDetails := make([]*user.PermDetail, 0)
	err = copier.Copy(&permDetails, res)
	if err != nil {
		return resp, err
	}

	total, err := l.svcCtx.PermModel.Count(l.ctx)
	if err != nil {
		return resp, err
	}
	return &user.ListPermResp{
		Page:       in.Page,
		PageSize:   in.PageSize,
		Total:      total,
		PermDetail: permDetails,
	}, nil
}
