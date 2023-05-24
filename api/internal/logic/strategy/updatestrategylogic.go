package strategy

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/rpc/userclient"

	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStrategyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStrategyLogic {
	return &UpdateStrategyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStrategyLogic) UpdateStrategy(req *types.UpdateStrategyReq) (resp *types.UpdateStrategyResp, err error) {
	var updateInfo = &userclient.StrategyInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.UpdateStrategy(l.ctx, &userclient.UpdateStrategyReq{
		Id:           req.Id,
		StrategyInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
