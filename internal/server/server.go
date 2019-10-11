package server

import(
	"github.com/gin-gonic/gin"
)

type Server struct {
	Route *gin.Engine
}

func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)
	return  &Server {
		Route: gin.Default(),
	}
}

func (srv *Server) Start() {

	srv.Route.Run()
}
