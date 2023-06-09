package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/rpc/client/roleservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleLogic) GetRole(req *types.GetRoleReq) (resp *types.GetRoleResp, err error) {
	res, err := l.svcCtx.RoleRpc.GetRoleById(l.ctx, &roleservice.GetRoleByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return resp, err
	}
	var roleDetail types.RoleDetail
	copier.Copy(&roleDetail, res.RoleDetail)
	resp = &types.GetRoleResp{RoleDetail: roleDetail}
	return
}
