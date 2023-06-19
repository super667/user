package auth

import (
	"github.com/super667/user/api/internal/logic/auth"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogoutReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout(&req)
		response.Response(w, resp, err)
	}
}
