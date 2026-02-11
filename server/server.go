package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hmdnu/fintr/internal/auth"
	"github.com/hmdnu/fintr/internal/user"
	"github.com/hmdnu/fintr/middleware"
)

type Server struct {
	User *user.Handler
	Auth *auth.Handler
}

func New(h *Server) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", h.Auth.Login)
	mux.HandleFunc("GET /logout", middleware.Auth(h.Auth.Logout))

	mux.HandleFunc("GET /user", middleware.Auth(h.User.List))
	mux.HandleFunc("POST /user", middleware.Auth(h.User.Create))

	mux.HandleFunc("POST /transaction", middleware.Auth(func(w http.ResponseWriter, r *http.Request) {}))
	return mux
}

func Listen(port string, mux *http.ServeMux) {
	fmt.Println("Server listening on port " + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
