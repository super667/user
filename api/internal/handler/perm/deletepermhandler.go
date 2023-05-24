package perm

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/perm"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func DeletePermHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePermReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := perm.NewDeletePermLogic(r.Context(), svcCtx)
		resp, err := l.DeletePerm(&req)
		response.Response(w, resp, err)
	}
}
