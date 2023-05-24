package userrole

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/userrole"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func UpdateUserRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userrole.NewUpdateUserRoleLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserRole(&req)
		response.Response(w, resp, err)
	}
}
