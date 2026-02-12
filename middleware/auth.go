package middleware

import (
	"net/http"

	"github.com/hmdnu/fintr/pkg/response"
	"github.com/hmdnu/fintr/pkg/token"
)

func Auth(handler AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			response.UnauthorizedErr(w, "token not found")
			return err
		}
		_, isTokenVerified := token.VerifiyToken(cookie.Value)
		if !isTokenVerified {
			response.UnauthorizedErr(w, "token not valid or expired")
			return err
		}
		return handler(w, r)
	}
}
