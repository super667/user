package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"user/common/errorx"
	"user/model"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermByIdLogic {
	return &GetPermByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPermByIdLogic) GetPermById(in *user.GetPermByIdReq) (*user.GetPermByIdResp, error) {
	Info, err := l.svcCtx.PermModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("权限点不存在")
	default:
		return nil, err
	}
	var permDetail user.PermDetail
	copier.Copy(&permDetail, Info)

	return &user.GetPermByIdResp{PermDetail: &permDetail}, nil
}
