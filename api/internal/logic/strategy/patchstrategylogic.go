package strategy

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/rpc/client/strategyservice"

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
	var updateInfo = &strategyservice.StrategyInfo{}
	copier.Copy(updateInfo, req)
	_, err = l.svcCtx.StrategyRpc.PatchStrategy(l.ctx, &strategyservice.PatchStrategyReq{
		Id:           req.Id,
		StrategyInfo: updateInfo,
	})
	if err != nil {
		return
	}
	return
}
