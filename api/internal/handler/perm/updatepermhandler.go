package perm

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/perm"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func UpdatePermHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePermReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := perm.NewUpdatePermLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePerm(&req)
		response.Response(w, resp, err)
	}
}
