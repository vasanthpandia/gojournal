package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/vasanthpandia/gojournal/internal/controllers"
)

func Login(c *gin.Context) {
	logger := c.MustGet("Logger").(*zap.Logger)
	controller := c.MustGet("SessionsController").(*controllers.SessionsController)
	request := &controllers.LoginPayload{}
	err := c.BindJSON(&request)

	if err != nil {
		logger.Error("Bind Error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	tokenstr, err := controller.Login(request)

	if err != nil {
		logger.Error("Bind Error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token" : tokenstr,
	})

}
