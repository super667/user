package middleware

import (
	"context"
	"github.com/super667/user/user/rpc/client/authservice"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
)

type TokenBlackListMiddleware struct {
	AuthRpc authservice.AuthService
}

func NewTokenBlackListMiddleware(t authservice.AuthService) *TokenBlackListMiddleware {
	return &TokenBlackListMiddleware{
		AuthRpc: t,
	}
}

func (m *TokenBlackListMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}

func TokenBlackList(s authservice.AuthService) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Middleware", "TokenBlackList")
			author := r.Header.Get("Authorization")
			token := strings.TrimPrefix(author, "Bearer ")
			resp, err := s.BlackListToken(r.Context(), &authservice.BlackListTokenReq{
				Token: token,
			})

			if err != nil {
				return
			}
			if resp.Invalid {
				return
			}

			ctx := context.WithValue(r.Context(), "token", token)
			next(w, r.WithContext(ctx))
		}
	}
}
