package permservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermLogic {
	return &CreatePermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePermLogic) CreatePerm(in *user.CreatePermReq) (*user.CreatePermResp, error) {
	resp := &user.CreatePermResp{}
	permInfo := &model.Permission{}
	err := copier.Copy(permInfo, in.PermInfo)
	if err != nil {
		l.Logger.Error(err)
		return resp, err
	}
	insertRes, err := l.svcCtx.PermModel.Insert(l.ctx, permInfo)
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
