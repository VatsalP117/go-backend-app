package main

import (
	"log"

	"github.com/VatsalP117/go-backend-app/internal/config"
	"github.com/VatsalP117/go-backend-app/internal/database"
	"github.com/VatsalP117/go-backend-app/internal/handlers" // We will create this next
	"github.com/VatsalP117/go-backend-app/internal/middleware"
	"github.com/VatsalP117/go-backend-app/internal/server"
)

func main() {
	cfg := config.Load()

	// 1. Initialize Database
	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	// Defer closing the connection until the app stops
	defer db.Close()
	log.Println("Connected to Database!")

	srv := server.NewServer(cfg)
	authMiddleware := middleware.New()
	
	// 2. Pass the DB to the Handler
	// We need to update NewUserHandler to accept the DB (see Step 6)
	userHandler := handlers.NewUserHandler(db) 

	api := srv.Echo.Group("/api")
	protected := api.Group("/v1")
	protected.Use(authMiddleware.RequireAuth)
	protected.GET("/profile", userHandler.GetProfile)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}