package server

import (
	"github.com/gin-gonic/gin"
	"github.com/symonk/toodoo/internal/controller"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	// Register default (unprotected, non versioned) routes
	healthCheckHandler := controller.HealthCheckHandler{}
	router.GET("/healthcheck", healthCheckHandler.Status)

	// Manage API versioned routes
	apiV1 := router.Group("/v1")
	apiV1.GET("/placeholder", func(c *gin.Context) {})

	return router
}
