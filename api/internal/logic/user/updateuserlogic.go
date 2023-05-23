package user

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	var updateInfo = &userclient.UserInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &userclient.UpdateUserReq{
		Id:       req.Id,
		UserInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
