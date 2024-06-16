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

	// Register default (unprotected, non versioned) routes
	healthCheckHandler := controller.HealthCheckHandler{}
	router.GET("/healthcheck", healthCheckHandler.Status)

	// Manage API versioned routes
	apiV1 := router.Group("/v1")
	apiV1.GET("/placeholder", func(c *gin.Context) {})

	return router
}
