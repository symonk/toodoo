package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var started = time.Now()

type HealthCheckHandler struct {
}

func (h HealthCheckHandler) Status(c *gin.Context) {
	since := time.Since(started)
	c.JSON(http.StatusOK, gin.H{"Uptime": since.String()})
}

// since returns the total uptime for the server in a
// string format.
func (h HealthCheckHandler) Since() string {
	return time.Since(started).String()
}
