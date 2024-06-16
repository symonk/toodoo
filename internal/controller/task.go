package controller

import "github.com/gin-gonic/gin"

type TaskHandler struct {
}

// @BasePage /api/v1
// Retrieve godoc
// @Summary Fetch all tasks
// @Schemes
// @Description Retrieves all the tasks
// @Tags tasks
// @Produce json
// @Success 200 {object} model.Task
// @Router /tasks/ [get]
func (t TaskHandler) View(c *gin.Context) {
	task := make(map[string]string)
	task["first task"] = "A first task"
	c.JSON(200, task)
}
