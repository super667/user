package menu

import (
	"context"

	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMenuLogic {
	return &ListMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMenuLogic) ListMenu(req *types.ListMenuReq) (resp *types.ListMenuResp, err error) {
	resp = &types.ListMenuResp{}
	resp.Menus = []types.MenuDetail{
		{
			Id: 1,
			MenuInfo: types.MenuInfo{
				AuthName: "用户管理",
				Path:     "users",
				Order:    0,
			},
			Children: []types.MenuDetail{
				{
					Id: 10,
					MenuInfo: types.MenuInfo{
						AuthName: "用户列表",
						Path:     "users",
						Order:    4,
					},
					Children: []types.MenuDetail{},
				},
			},
		},
		{
			Id: 2,
			MenuInfo: types.MenuInfo{
				AuthName: "权限管理",
				Path:     "rights",
				Order:    5,
			},
			Children: []types.MenuDetail{
				{
					Id: 10,
					MenuInfo: types.MenuInfo{
						AuthName: "权限列表",
						Path:     "rights",
						Order:    4,
					},
					Children: []types.MenuDetail{},
				},
				{
					Id: 10,
					MenuInfo: types.MenuInfo{
						AuthName: "角色列表",
						Path:     "roles",
						Order:    4,
					},
					Children: []types.MenuDetail{},
				},
			},
		},
		{
			Id: 3,
			MenuInfo: types.MenuInfo{
				AuthName: "商品管理",
				Path:     "goods",
				Order:    1,
			},
			Children: []types.MenuDetail{
				{
					Id: 10,
					MenuInfo: types.MenuInfo{
						AuthName: "商品列表",
						Path:     "goods",
						Order:    4,
					},
					Children: []types.MenuDetail{},
				},
			},
		},
		{
			Id: 4,
			MenuInfo: types.MenuInfo{
				AuthName: "订单管理",
				Path:     "orders",
				Order:    2,
			},
			Children: []types.MenuDetail{
				{
					Id: 10,
					MenuInfo: types.MenuInfo{
						AuthName: "订单列表",
						Path:     "orders",
						Order:    4,
					},
					Children: []types.MenuDetail{},
				},
			},
		},
		{
			Id: 5,
			MenuInfo: types.MenuInfo{
				AuthName: "数据统计",
				Path:     "reports",
				Order:    3,
			},
			Children: []types.MenuDetail{
				{
					Id: 10,
					MenuInfo: types.MenuInfo{
						AuthName: "统计列表",
						Path:     "reports",
						Order:    4,
					},
					Children: []types.MenuDetail{},
				},
			},
		},
	}
	return
}
