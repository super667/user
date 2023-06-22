package userrole

import (
	"github.com/super667/user/user/api/internal/logic/userrole"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CreateUserRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userrole.NewCreateUserRoleLogic(r.Context(), svcCtx)
		resp, err := l.CreateUserRole(&req)
		response.Response(w, resp, err)
	}
}
