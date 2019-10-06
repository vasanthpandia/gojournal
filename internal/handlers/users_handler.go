package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/vasanthpandia/gojournal/internal/controllers"
)

func CreateUser(c *gin.Context) {
	controller := c.MustGet("UsersController").(*controllers.UsersController)

	var request controllers.UserCreatePayload

	err := c.BindJSON(request)

	if err != nil {
		c.String(http.StatusInternalServerError, "Bind Failed")
	}

	user, err := controller.Create(&request)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, user)
}
