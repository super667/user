package strategyservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/common/errorx"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStrategyByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStrategyByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStrategyByIdLogic {
	return &GetStrategyByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStrategyByIdLogic) GetStrategyById(in *user.GetStrategyByIdReq) (*user.GetStrategyByIdResp, error) {
	Info, err := l.svcCtx.StrategyModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("赋权策略不存在")
	default:
		return nil, err
	}
	var strategyDetail user.StrategyDetail
	copier.Copy(&strategyDetail, Info)

	return &user.GetStrategyByIdResp{StrategyDetail: &strategyDetail}, nil
}
