package main

import (
	"context"
	"embed"
	"fmt"
	"goweb/app"
	"goweb/controller"
	"goweb/middleware"
	"goweb/repository"
	"goweb/service"
	"goweb/utils"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//go:embed views/* static/*
var content embed.FS

func main() {
	mux := http.DefaultServeMux

	// mux.Handle("/static/",
	// 	http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
	// )
	//

	// ENV
	serverPort := os.Getenv("SERVER_PORT")
	dbConnection := os.Getenv("PSQL")
	//

	db := app.GetDatabaseConnection(dbConnection)
	// Close DB Pool before Serving
	defer db.Close()

	validatorInstance := utils.GetValidator()
	var handler http.Handler = mux

	// Dependency Injection
	authRepo := repository.NewAuthRepository()
	authService := service.NewAuthService(db, authRepo, validatorInstance)
	authController := controller.NewAuthController(authService)

	todoRepo := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepo, validatorInstance)
	todoController := controller.NewTodoController(todoService)

	staticPageController := controller.NewStaticPageController(content)

	// Router setup
	router := app.Router(authController, todoController, staticPageController)
	handler = middleware.Middleware(router)

	//
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	/// Server Setup
	if serverPort == "" {
		serverPort = "9002"
	}
	svr := new(http.Server)
	svr.Addr = ":" + serverPort
	svr.Handler = handler

	go func() {
		fmt.Printf("Server started at port %s\n", svr.Addr)
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server failed: ", err.Error())
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutting down...")

	shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := svr.Shutdown(shutDownCtx); err != nil {
		fmt.Println("Shutdown failed: ", err)
	}

	fmt.Println("Server is shutdown")

}
