package controller

import (
	"fmt"
	"net/http"
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
// @Router /task [get]
func (t TaskHandler) View(c *gin.Context) {
	tasks, err := taskModel.RetrieveTasks(c.Request.Context())
	if err != nil {
		// TODO: standardize and improve errors in general etc.
		// Do not propagate internal errors back to clients;
		// Doing this for development only for now;
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, tasks)
}

// @BasePage /api/v1
// Retrieve godoc
// @Summary Fetch all tasks
// @Schemes
// @Description Retrieves all the tasks
// @Tags tasks
// @Produce json
// @Success 200 {object} model.TaskModel
// @Router /task/id [get]
func (t TaskHandler) ViewByID(c *gin.Context) {
	id := c.Param("id")
	asInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be a valid integer")
	}
	task, err := taskModel.RetrieveTaskByID(c.Request.Context(), asInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, task)
}

// @BasePage /api/v1
// Retrieve godoc
// @Summary Creates a new task
// @Schemes
// @Description Creates a new task
// @Tags tasks
// @Produce json
// @Accept json
// @Success 201 {object} model.TaskModel
// @Router /task [post]
func (t TaskHandler) Create(c *gin.Context) {
	var newTask model.TaskModel
	if err := c.ShouldBind(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The task creation payload was not valid: %s", err.Error())})
		return
	}
	serverSideTask, err := taskModel.Create(c.Request.Context(), newTask)
	if err != nil {
		// TODO: hide error, only debugging for now
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, serverSideTask)
}
