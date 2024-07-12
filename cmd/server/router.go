package server

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	docs "github.com/symonk/toodoo/docs"
	"github.com/symonk/toodoo/internal/controller"
	"github.com/symonk/toodoo/internal/logging"
	"github.com/symonk/toodoo/internal/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(sloggin.New(logging.Logger))
	// TODO: Have this properly handle all scenarios
	//router.Use(middleware.JsonIndenter)

	// Handlers
	healthCheckHandler := controller.HealthCheckHandler{}
	taskHandler := controller.TaskHandler{}
	userHandler := controller.Userhandler{}
	authhandler := controller.AuthHandler{}

	// Register default (unprotected, non versioned) routes
	router.GET("/healthcheck", healthCheckHandler.Status)

	// Manage API versioned routes
	// protection is currently being implemented here.
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/validate", middleware.Protected, authhandler.Validate)

	// task management
	apiV1.GET("/task", taskHandler.View)
	apiV1.GET("/task/:id", taskHandler.ViewByID)
	apiV1.POST("/task", taskHandler.Create)

	// user management
	apiV1.GET("/user", userHandler.View)

	// Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
