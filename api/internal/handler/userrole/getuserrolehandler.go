package userrole

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/userrole"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func GetUserRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userrole.NewGetUserRoleLogic(r.Context(), svcCtx)
		resp, err := l.GetUserRole(&req)
		response.Response(w, resp, err)
	}
}
