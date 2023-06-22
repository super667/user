package user

import (
	"github.com/super667/user/user/api/response"
	"net/http"

	"github.com/super667/user/user/api/internal/logic/user"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PatchUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PatchUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewPatchUserLogic(r.Context(), svcCtx)
		resp, err := l.PatchUser(&req)
		response.Response(w, resp, err)
	}
}
