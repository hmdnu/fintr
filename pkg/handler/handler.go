package handler

import (
	"net/http"

	"github.com/hmdnu/fintr/middleware"
)

type HandlerRoute struct {
	mux        *http.ServeMux
	middleware []middleware.Middleware
}

func NewRoute(middleware ...middleware.Middleware) *HandlerRoute {
	return &HandlerRoute{mux: http.NewServeMux(), middleware: middleware}
}

func (r *HandlerRoute) Handle(pattern string, h middleware.AppHandler, m ...middleware.Middleware) {
	all := append(m, r.middleware...)
	chainedMiddleware := middleware.ChainMiddleware(h, all...)
	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		chainedMiddleware(w, r)
	})
}

func (r *HandlerRoute) GetMux() *http.ServeMux {
	return r.mux
}
