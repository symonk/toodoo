package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/symonk/toodoo/internal/model"
)

type TaskHandler struct {
}

var taskModel = new(model.TaskModel)

// @BasePage /api/v1
// Retrieve godoc
// @Summary Fetch all tasks
// @Schemes
// @Description Retrieves all the tasks
// @Tags tasks
// @Produce json
// @Success 200 {object} []model.TaskModel
// @Router /tasks/ [get]
func (t TaskHandler) View(c *gin.Context) {
	tasks, err := taskModel.RetrieveTasks(c.Request.Context())
	if err != nil {
		// TODO: standardize and improve errors in general etc.
		// Do not propagate internal errors back to clients;
		// Doing this for development only for now;
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, tasks)
}

// @BasePage /api/v1
// Retrieve godoc
// @Summary Fetch all tasks
// @Schemes
// @Description Retrieves all the tasks
// @Tags tasks
// @Produce json
// @Success 200 {object} model.TaskModel
// @Router /tasks/id [get]
func (t TaskHandler) ViewByID(c *gin.Context) {
	id := c.Param("id")
	asInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(400, "id must be a valid integer")
	}
	task, err := taskModel.RetrieveTaskByID(c.Request.Context(), asInt)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, task)
}
