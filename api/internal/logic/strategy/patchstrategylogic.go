package strategy

import (
	"context"
	"github.com/jinzhu/copier"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchStrategyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchStrategyLogic {
	return &PatchStrategyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchStrategyLogic) PatchStrategy(req *types.PatchStrategyReq) (resp *types.PatchStrategyResp, err error) {
	var updateInfo = &userclient.StrategyInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.UserRpc.PatchStrategy(l.ctx, &userclient.PatchStrategyReq{
		Id:           req.Id,
		StrategyInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}