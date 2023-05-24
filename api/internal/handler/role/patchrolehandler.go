package role

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/role"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func PatchRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PatchRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewPatchRoleLogic(r.Context(), svcCtx)
		resp, err := l.PatchRole(&req)
		response.Response(w, resp, err)
	}
}
