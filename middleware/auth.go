package middleware

import (
	"net/http"

	"github.com/hmdnu/fintr/pkg/response"
	"github.com/hmdnu/fintr/pkg/token"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			response.UnauthorizedErr(w, "token not found")
			return
		}
		_, isTokenVerified := token.VerifiyToken(cookie.Value)
		if !isTokenVerified {
			response.UnauthorizedErr(w, "token not valid or expired")
			return
		}
		handler.ServeHTTP(w, r)
	})

}
