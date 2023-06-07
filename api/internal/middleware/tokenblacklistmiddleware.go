package middleware

import (
	"github.com/super667/user/model"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

type TokenBlackListMiddleware struct {
	tokenModel model.TokenModel
}

func NewTokenBlackListMiddleware(t model.TokenModel) *TokenBlackListMiddleware {
	return &TokenBlackListMiddleware{
		tokenModel: t,
	}
}

func (m *TokenBlackListMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}

func TokenBlackList(s model.TokenModel) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			_, _ = s.FindOne(r.Context(), 0)
			w.Header().Add("X-Middleware", "TokenBlackList")
			next(w, r)
		}
	}
}
