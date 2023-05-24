package strategy

import (
	"github.com/super667/user/api/internal/logic/strategy"
	"github.com/super667/user/api/internal/svc"
	"github.com/super667/user/api/internal/types"
	"github.com/super667/user/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UpdateStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateStrategyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := strategy.NewUpdateStrategyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStrategy(&req)
		response.Response(w, resp, err)
	}
}
