package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"user/model"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleLogic) CreateRole(in *user.CreateRoleReq) (*user.CreateRoleResp, error) {
	resp := &user.CreateRoleResp{}
	roleInfo := &model.Role{}
	err := copier.Copy(roleInfo, in.RoleInfo)
	if err != nil {
		l.Logger.Error(err)
		return resp, err
	}
	insertRes, err := l.svcCtx.RoleModel.Insert(l.ctx, roleInfo)
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
