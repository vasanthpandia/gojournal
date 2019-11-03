package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/vasanthpandia/gojournal/internal/controllers"
	"github.com/vasanthpandia/gojournal/internal/jsonerrors"
)

func Login(c *gin.Context) {
	logger := c.MustGet("Logger").(*zap.Logger)
	controller := c.MustGet("SessionsController").(*controllers.SessionsController)
	request := &controllers.LoginPayload{}
	err := c.BindJSON(&request)

	if err != nil {
		logger.Error("Bind Error", zap.Error(err))
		c.JSON(http.StatusBadRequest, jsonerrors.New("Bad Request Body"))
	}

	authtoken, err := controller.Login(request)

	if err != nil {
		logger.Error("Bind Error", zap.Error(err))
		c.JSON(http.StatusBadRequest, jsonerrors.New("Bad Request Body"))
	}

	c.JSON(http.StatusOK, authtoken)
}
