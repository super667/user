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

type GetRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIdLogic {
	return &GetRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleByIdLogic) GetRoleById(in *user.GetRoleByIdReq) (*user.GetRoleByIdResp, error) {
	Info, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("角色不存在")
	default:
		return nil, err
	}
	var roleDetail user.RoleDetail
	copier.Copy(&roleDetail, Info)

	return &user.GetRoleByIdResp{RoleDetail: &roleDetail}, nil
}
