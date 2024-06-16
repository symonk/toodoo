package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

var started = time.Now()

type Handler struct{}

type HealthCheckHandler struct {
}

func (h HealthCheckHandler) Status(c *gin.Context) {
	since := time.Since(started)
	c.JSON(200, gin.H{"Uptime": since.String()})
}

// since returns the total uptime for the server in a
// string format.
func (h HealthCheckHandler) Since() string {
	return time.Since(started).String()
}
