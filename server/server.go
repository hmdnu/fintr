package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hmdnu/fintr/internal/auth"
	"github.com/hmdnu/fintr/internal/category"
	"github.com/hmdnu/fintr/internal/transaction"
	"github.com/hmdnu/fintr/internal/user"
	"github.com/hmdnu/fintr/middleware"
	"github.com/hmdnu/fintr/pkg/handler"
)

type Server struct {
	User        *user.Handler
	Auth        *auth.Handler
	Transaction *transaction.Handler
	Category    *category.Handler
}

func New(h *Server) *http.ServeMux {
	route := handler.NewRoute(middleware.Logger)

	route.Handle("POST /login", h.Auth.Login)
	route.Handle("GET /logout", h.Auth.Logout, middleware.Auth)
	route.Handle("GET /user", h.User.List, middleware.Auth)
	route.Handle("POST /user", h.User.Create, middleware.Auth)
	route.Handle("GET /transaction", h.Transaction.List, middleware.Auth)
	route.Handle("POST /transaction", h.Transaction.Create, middleware.Auth)
	route.Handle("POST /category", h.Category.Create, middleware.Auth)
	route.Handle("GET /category", h.Category.List, middleware.Auth)

	return route.GetMux()
}

func Listen(port string, mux *http.ServeMux) {
	fmt.Println("Server listening on port " + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
