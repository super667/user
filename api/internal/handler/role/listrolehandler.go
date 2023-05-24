package role

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/role"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func ListRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewListRoleLogic(r.Context(), svcCtx)
		resp, err := l.ListRole(&req)
		response.Response(w, resp, err)
	}
}
