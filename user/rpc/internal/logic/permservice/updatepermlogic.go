package permservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/model"

	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermLogic {
	return &UpdatePermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePermLogic) UpdatePerm(in *user.UpdatePermReq) (*user.UpdatePermResp, error) {
	var permInfo model.Permission
	copier.Copy(&permInfo, in.PermInfo)
	permInfo.Id = in.Id
	err := l.svcCtx.PermModel.Update(l.ctx, &permInfo)
	if err != nil {
		return &user.UpdatePermResp{}, err
	}

	return &user.UpdatePermResp{}, nil
}
