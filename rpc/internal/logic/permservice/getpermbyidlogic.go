package permservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/common/errorx"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

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
