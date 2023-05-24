package perm

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/rpc/userclient"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

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
	var updateInfo = &userclient.PermInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.PatchPerm(l.ctx, &userclient.PatchPermReq{
		Id:       req.Id,
		PermInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
