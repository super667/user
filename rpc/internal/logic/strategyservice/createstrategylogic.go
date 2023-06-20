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
	subjectType := in.StrategyInfo.SubjectType
	perm := in.StrategyInfo.Perm
	var strategyInfos = make([]*model.Strategy, 0, len(in.StrategyInfo.Subject)*len(in.StrategyInfo.Resource))
	for _, s := range in.StrategyInfo.Subject {
		for _, r := range in.StrategyInfo.Resource {
			strategyInfos = append(strategyInfos, &model.Strategy{
				Subject:     s,
				SubjectType: subjectType,
				Resource:    r,
				Perm:        perm,
			})
		}
	}

	err = l.svcCtx.StrategyModel.TranBatchInsert(l.ctx, strategyInfos)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
