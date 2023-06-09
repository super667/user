package strategyservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStrategyLogic {
	return &CreateStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStrategyLogic) CreateStrategy(in *user.CreateStrategyReq) (*user.CreateStrategyResp, error) {
	resp := &user.CreateStrategyResp{}
	strategyInfo := &model.Strategy{}
	err := copier.Copy(strategyInfo, in.StrategyInfo)
	if err != nil {
		l.Logger.Error(err)
		return resp, err
	}
	insertRes, err := l.svcCtx.StrategyModel.Insert(l.ctx, strategyInfo)
	if err != nil {
		return resp, err
	}
	lastId, err := insertRes.LastInsertId()
	if err != nil {
		return resp, err
	}
	resp.Id = lastId

	return resp, nil
}
