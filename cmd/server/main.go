package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ardnh/be-travel-booking-app/internal/config"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/database/postgresql"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/database/redis"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/middleware"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/routes"
	logger "github.com/ardnh/be-travel-booking-app/internal/utils/logger"
	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3" // tetap sama!
	"github.com/gofiber/fiber/v3"
)

func main() {

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory:", err)
	}

	modelPath := filepath.Join(workDir, "internal/config", "casbin_model.conf")

	// TODO: Implement application initialization
	// 1. Load configuration
	logger := logger.New()
	cfg := config.LoadConfig()
	// validator := validator.New()

	// 2. Initialize database
	db, err := postgresql.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer postgresql.CloseDB(db)

	redisDb := redis.NewRedisDB(cfg)
	defer redisDb.Close()

	// 4. Wire up dependencies
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatal(err)
	}

	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		log.Fatal(err)
	}
	enforcer.LoadPolicy()

	// 5. Start HTTP server
	app := fiber.New()

	requestTimer := middleware.NewRequestTimerMiddleware(logger)
	app.Use(requestTimer.Track())

	// Repository
	// Service
	// Handler

	routes.SetupAPIRoutes(
		app,
		logger,
		enforcer,
	)

	portListen := fmt.Sprintf(":%s", cfg.App.Port)
	if err := app.Listen(portListen); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
	os.Exit(0)
}
