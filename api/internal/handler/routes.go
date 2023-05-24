// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	perm "user/api/internal/handler/perm"
	role "user/api/internal/handler/role"
	strategy "user/api/internal/handler/strategy"
	user "user/api/internal/handler/user"
	userrole "user/api/internal/handler/userrole"
	"user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/:id",
				Handler: user.GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user",
				Handler: user.CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user",
				Handler: user.ListUserHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/:id",
				Handler: user.DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/:id",
				Handler: user.UpdateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/user/:id",
				Handler: user.PatchUserHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/role/:id",
				Handler: role.GetRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role",
				Handler: role.CreateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role",
				Handler: role.ListRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/role/:id",
				Handler: role.DeleteRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/role/:id",
				Handler: role.UpdateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/role/:id",
				Handler: role.PatchRoleHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/perm/:id",
				Handler: perm.GetPermHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/perm",
				Handler: perm.CreatePermHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/perm",
				Handler: perm.ListPermHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/perm/:id",
				Handler: perm.DeletePermHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/perm/:id",
				Handler: perm.UpdatePermHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/perm/:id",
				Handler: perm.PatchPermHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/strategy/:id",
				Handler: strategy.GetStrategyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/strategy",
				Handler: strategy.CreateStrategyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/strategy",
				Handler: strategy.ListStrategyHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/strategy/:id",
				Handler: strategy.DeleteStrategyHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/strategy/:id",
				Handler: strategy.UpdateStrategyHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/strategy/:id",
				Handler: strategy.PatchStrategyHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user_role/:id",
				Handler: userrole.GetUserRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user_role",
				Handler: userrole.CreateUserRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user_role",
				Handler: userrole.ListUserRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user_role/:id",
				Handler: userrole.DeleteUserRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user_role/:id",
				Handler: userrole.UpdateUserRoleHandler(serverCtx),
			},
		},
	)
}
