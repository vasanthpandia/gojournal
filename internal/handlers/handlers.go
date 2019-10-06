package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/vasanthpandia/gojournal/internal/controllers"
)

// type Payload struct {
// }

func BasicHandler(c *gin.Context) *gin.Handler {
	response := controllers.BasicController()

	c.String(http.StatusOK, response)
}
