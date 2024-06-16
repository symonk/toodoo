package server

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"github.com/symonk/toodoo/internal/controller"
	"github.com/symonk/toodoo/internal/logging"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(sloggin.New(logging.Logger))

	// Handlers
	healthCheckHandler := controller.HealthCheckHandler{}
	taskHandler := controller.TaskHandler{}

	// Register default (unprotected, non versioned) routes
	router.GET("/healthcheck", healthCheckHandler.Status)

	// Manage API versioned routes
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/task", taskHandler.View)

	return router
}
