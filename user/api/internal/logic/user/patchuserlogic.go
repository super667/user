package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/userservice"

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
	var updateInfo = &userservice.UserInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.PatchUser(l.ctx, &userservice.PatchUserReq{
		Id:       req.Id,
		UserInfo: updateInfo,
	})
	if err != nil {
		return
	}

	return
}
