package user

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchUserLogic {
	return &PatchUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchUserLogic) PatchUser(req *types.PatchUserReq) (resp *types.PatchUserResp, err error) {
	var updateInfo = &userclient.UserInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.PatchUser(l.ctx, &userclient.PatchUserReq{
		Id:       req.Id,
		UserInfo: updateInfo,
	})
	if err != nil {
		return
	}

	return
}
