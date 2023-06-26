package middleware

import (
	"github.com/super667/user/user/model"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

type EmailMiddleware struct {
}

func NewEmailMiddleware() *EmailMiddleware {
	return &EmailMiddleware{}
}

func (m *EmailMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}

func EmailWithUser(s model.UserModel) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			_, _ = s.GetOne(r.Context(), "id")
			w.Header().Add("X-Middleware", "demo")
			next(w, r)
		}
	}
}
