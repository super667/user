package perm

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/permservice"

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
	var updateInfo = &permservice.PermInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.PermRpc.UpdatePerm(l.ctx, &permservice.UpdatePermReq{
		Id:       req.Id,
		PermInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
