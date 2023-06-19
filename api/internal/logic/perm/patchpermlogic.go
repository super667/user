package perm

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/permservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchPermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchPermLogic {
	return &PatchPermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchPermLogic) PatchPerm(req *types.PatchPermReq) (resp *types.PatchPermResp, err error) {
	var updateInfo = &permservice.PermInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.PermRpc.PatchPerm(l.ctx, &permservice.PatchPermReq{
		Id:       req.Id,
		PermInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
