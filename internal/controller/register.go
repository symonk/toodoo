package controller

import "github.com/gin-gonic/gin"

type RegisterHandler struct {
}

// @BasePage /api/v1
// Retrieve godoc
// @Registers a new user to the system
// @Schemes
// @Description Registers a new user to the system
// @Tags register
// @Produce json
// @Success 201 {object} []model.TaskModel
// @Router /task [post]
func (r RegisterHandler) Register(c *gin.Context) {

}
