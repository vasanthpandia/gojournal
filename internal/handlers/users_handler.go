package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/vasanthpandia/gojournal/internal/controllers"
)

func CreateUser(c *gin.Context) {
	controller := c.MustGet("UsersController").(*controllers.UsersController)
	logger := c.MustGet("Logger").(*zap.Logger)

	var request *controllers.UserCreatePayload

	err := c.BindJSON(request)

	if err != nil {
		logger.Error("Json Bind Error", zap.Error(err))
		c.String(http.StatusInternalServerError, "Bind Failed")
	}

	user, err := controller.Create(request)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, user)
}
