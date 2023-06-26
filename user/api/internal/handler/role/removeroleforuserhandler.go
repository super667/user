package role

import (
	"github.com/super667/user/user/api/internal/logic/role"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/super667/user/user/api/internal/types"
	"github.com/super667/user/user/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func RemoveRoleForUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveRoleForUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRemoveRoleForUserLogic(r.Context(), svcCtx)
		resp, err := l.RemoveRoleForUser(&req)
		response.Response(w, resp, err)
	}
}
