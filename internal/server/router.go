package server

import "github.com/gin-gonic/gin"

func newRouter() *gin.Engine {
	router := gin.Default()
	return router
}
