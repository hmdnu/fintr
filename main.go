package main

import (
	"log"

	"github.com/hmdnu/fintr/internal/auth"
	"github.com/hmdnu/fintr/internal/category"
	"github.com/hmdnu/fintr/internal/transaction"
	"github.com/hmdnu/fintr/internal/user"
	"github.com/hmdnu/fintr/pkg/database"
	_ "github.com/hmdnu/fintr/pkg/database"
	"github.com/hmdnu/fintr/pkg/env"
	"github.com/hmdnu/fintr/server"
)

func main() {
	env.Load()
	conn, err := database.Connect()
	// database.InitTableIfNotExist(conn)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	authService := auth.NewService(conn)
	authHandler := auth.NewHandler(authService)
	userService := user.NewService(conn)
	userHandler := user.NewHandler(userService)
	transactionServie := transaction.NewServie(conn)
	transactionHandler := transaction.NewHandler(transactionServie)
	categoryService := category.NewService(conn)
	categoryHandler := category.NewHandler(categoryService)

	mux := server.New(&server.Server{User: userHandler, Auth: authHandler, Transaction: transactionHandler, Category: categoryHandler})
	server.Listen("8080", mux)
}
