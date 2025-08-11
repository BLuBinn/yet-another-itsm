package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"msn-map-api/internal/config"
	"msn-map-api/internal/controller"
	"msn-map-api/internal/database"
	"msn-map-api/internal/middleware"
	"msn-map-api/internal/repository"
	"msn-map-api/internal/router"
	"msn-map-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	log.Info().
		Str("port", cfg.Server.Port).
		Str("log_level", cfg.Logger.Level).
		Str("log_format", cfg.Logger.Format).
		Msg("Starting MSN Map API server")

	// Initialize database
	db, err := database.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}
	defer db.Close()

	// Initialize repository
	repository := repository.New(db.Pool)

	// Initialize services
	services := service.NewServices(db, repository, cfg)

	// Initialize controllers
	controllers := controller.NewControllers(services)

	// Setup Gin router
	if cfg.Logger.Level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin engine
	ginEngine := gin.New()

	// Add middleware
	ginEngine.Use(middleware.ZerologMiddleware())
	ginEngine.Use(middleware.RecoveryWithZerolog())
	ginEngine.Use(middleware.CORSMiddleware())
	ginEngine.Use(middleware.SecurityMiddleware())

	// Setup routes
	routers := router.NewRouters(controllers, cfg)
	routers.SetupRoutes(ginEngine)

	// Create HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      ginEngine,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		log.Info().
			Str("addr", server.Addr).
			Msg("Starting HTTP server")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}()

	log.Info().
		Str("version", "1.0.0").
		Str("environment", gin.Mode()).
		Msg("MSN Map API server started successfully")

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server forced to shutdown")
	} else {
		log.Info().Msg("Server exited gracefully")
	}
}
