package main

import (
	"game-hub-backend/internal/delivery/http/middleware"
	"log"

	"game-hub-backend/internal/di"
	"game-hub-backend/internal/infra/config"
	dbinfra "game-hub-backend/internal/infra/database"
	"game-hub-backend/internal/infra/database/migrations"
	"game-hub-backend/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect database
	db, err := dbinfra.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connected successfully")

	// Auto migrate
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")

	// Dependency Injection Container
	handlers := di.InitHandlers(db, cfg)

	// Router
	r := gin.Default()

	//Global Middleware
	middleware.SetupCors(r)

	// Register Routes
	router.RegisterRoutes(r, handlers)

	// Start server
	log.Println("Server running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
