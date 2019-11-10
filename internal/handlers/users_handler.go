package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/vasanthpandia/gojournal/internal/controllers"
	"github.com/vasanthpandia/gojournal/internal/models"
	"github.com/vasanthpandia/gojournal/internal/jsonerrors"
)

func CreateUser(c *gin.Context) {
	controller := c.MustGet("UsersController").(*controllers.UsersController)
	logger := c.MustGet("Logger").(*zap.Logger)

	request := &controllers.UserCreatePayload{}
	err := c.BindJSON(&request)

	if err != nil {
		logger.Error("Json Bind Error", zap.Error(err))
		c.JSON(http.StatusBadRequest, jsonerrors.BadRequest)
		return
	}

	user, err := controller.Create(request)

	if err != nil {
		logger.Error("Unknown Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, jsonerrors.New(err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	controller := c.MustGet("UsersController").(*controllers.UsersController)
	logger := c.MustGet("Logger").(*zap.Logger)

	userId := c.Param("userId")

	user, err := controller.Read(userId)

	if err != nil {
		logger.Error("Fetch User Failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, jsonerrors.New(err.Error()))
	}

	c.JSON(http.StatusOK, user)
}

func CurrentUser( c *gin.Context) {
	currentUser := c.MustGet("CurrentUser").(*models.User)

	c.JSON(http.StatusOK, currentUser)
}
