package strategy

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/strategy"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func DeleteStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteStrategyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := strategy.NewDeleteStrategyLogic(r.Context(), svcCtx)
		resp, err := l.DeleteStrategy(&req)
		response.Response(w, resp, err)
	}
}
