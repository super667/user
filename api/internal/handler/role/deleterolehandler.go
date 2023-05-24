package role

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/role"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func DeleteRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewDeleteRoleLogic(r.Context(), svcCtx)
		resp, err := l.DeleteRole(&req)
		response.Response(w, resp, err)
	}
}
