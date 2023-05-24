package strategy

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/api/internal/logic/strategy"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/response"
)

func CreateStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateStrategyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := strategy.NewCreateStrategyLogic(r.Context(), svcCtx)
		resp, err := l.CreateStrategy(&req)
		response.Response(w, resp, err)
	}
}
