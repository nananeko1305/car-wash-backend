package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/nananeko1305/car-wash-backend/internal/database"
	"github.com/nananeko1305/car-wash-backend/internal/db"
	"github.com/nananeko1305/car-wash-backend/internal/handler"
	"github.com/nananeko1305/car-wash-backend/internal/repository"
	"github.com/nananeko1305/car-wash-backend/internal/service"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	DBConnection, err := database.PGClient(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer DBConnection.Close()

	queries := db.New(DBConnection)

	// Define repositories
	userRepository := repository.InitUserRepository(queries)
	// Define services
	userService := service.InitUserService(userRepository)
	// Define handlers
	authHandler := handler.InitAuthHandler(userService)
	userHandler := handler.InitUserHandler(userService)

	rootRouter := chi.NewRouter()
	rootRouter.Use(middleware.Logger)
	rootRouter.Mount("/api", apiRouter(authHandler, userHandler))

	if err := http.ListenAndServe(":3000", rootRouter); err != nil {
		log.Fatal(err)
	}

}

func apiRouter(authHandler *handler.AuthHandler, userHandler *handler.UserHandler) http.Handler {

	apiRouter := chi.NewRouter()

	// AUTH
	apiRouter.Post("/login", authHandler.Login)

	// USERS
	apiRouter.Get("/users", userHandler.GetUsers)
	apiRouter.Post("/users", userHandler.CreateUser)

	return apiRouter
}
