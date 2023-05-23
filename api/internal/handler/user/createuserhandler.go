package user

import (
	"net/http"
	"user/api/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/api/internal/logic/user"
	"user/api/internal/svc"
	"user/api/internal/types"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		response.Response(w, resp, err)
	}
}
