package role

import (
	"github.com/super667/user/api/internal/logic/role"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ListUserRolesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListUserRolesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewListUserRolesLogic(r.Context(), svcCtx)
		resp, err := l.ListUserRoles(&req)
		response.Response(w, resp, err)
	}
}
