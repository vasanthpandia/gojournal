package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/vasanthpandia/gojournal/internal/controllers"
)

// type Payload struct {
// }

func BasicHandler(c *gin.Context) {
	response := controllers.BasicController()

	c.String(http.StatusOK, response)
}
