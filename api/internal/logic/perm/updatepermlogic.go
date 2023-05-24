package perm

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermLogic {
	return &UpdatePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePermLogic) UpdatePerm(req *types.UpdatePermReq) (resp *types.UpdatePermResp, err error) {
	var updateInfo = &userclient.PermInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.UpdatePerm(l.ctx, &userclient.UpdatePermReq{
		Id:       req.Id,
		PermInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
