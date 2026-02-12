package middleware

import "net/http"

type AppHandler func(http.ResponseWriter, *http.Request) error
type Middleware func(AppHandler) AppHandler

func ChainMiddleware(h AppHandler, m ...Middleware) AppHandler {
	for i := range m {
		h = m[i](h)
	}
	return h
}
