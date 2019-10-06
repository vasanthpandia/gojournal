package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/vasanthpandia/gojournal/internal/controllers"
)

func CreateUser(c *gin.Context) {
	controller := c.MustGet("UsersController")

	request := new(controllers.UserCreatePayload)

	user, err := controller.Create(request)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, user)
}
