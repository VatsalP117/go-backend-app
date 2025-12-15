package main

import (
	"log"

	"github.com/VatsalP117/go-backend-app/internal/config"
	"github.com/VatsalP117/go-backend-app/internal/handlers" // We will create this next
	"github.com/VatsalP117/go-backend-app/internal/middleware"
	"github.com/VatsalP117/go-backend-app/internal/server"
)

func main() {
	// 1. Load Configuration
	cfg := config.Load()

	// 2. Initialize Server (Echo + Clerk Global Key)
	srv := server.NewServer(cfg)

	// 3. Initialize Internal Components
	authMiddleware := middleware.New()
	userHandler := handlers.NewUserHandler()

	// 4. Register Routes
	// We group routes by prefix for cleaner URLs
	api := srv.Echo.Group("/api")

	// Protected Routes (Require Clerk Token)
	protected := api.Group("/v1")
	protected.Use(authMiddleware.RequireAuth) // <--- The Gatekeeper
	
	// Map the URL to the Handler Function
	protected.GET("/profile", userHandler.GetProfile)

	// 5. Start the Server
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
