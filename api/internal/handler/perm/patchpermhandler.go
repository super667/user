package perm

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/perm"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func PatchPermHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PatchPermReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := perm.NewPatchPermLogic(r.Context(), svcCtx)
		resp, err := l.PatchPerm(&req)
		response.Response(w, resp, err)
	}
}
