package role

import (
	"github.com/super667/user/api/internal/logic/role"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ListAllUserRolesReqHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListAllUserRolesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewListAllUserRolesReqLogic(r.Context(), svcCtx)
		resp, err := l.ListAllUserRolesReq(&req)
		response.Response(w, resp, err)
	}
}
