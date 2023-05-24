package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchPermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchPermLogic {
	return &PatchPermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchPermLogic) PatchPerm(in *user.PatchPermReq) (*user.PatchPermResp, error) {
	var permInfo model.Permission
	copier.Copy(&permInfo, in.PermInfo)
	permInfo.Id = in.Id
	err := l.svcCtx.PermModel.PartialUpdate(l.ctx, &permInfo)
	if err != nil {
		return &user.PatchPermResp{}, err
	}

	return &user.PatchPermResp{}, nil
}
