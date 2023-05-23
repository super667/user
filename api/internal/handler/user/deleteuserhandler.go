package user

import (
	"net/http"
	"user/api/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/api/internal/logic/user"
	"user/api/internal/svc"
	"user/api/internal/types"
)

func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDeleteUserLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUser(&req)
		response.Response(w, resp, err)
	}
}
