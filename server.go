package server

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	engine *gin.Engine
	handlers []*gin.Handler
}

func GetDefaultServer() *gin.Engine {
	return gin.Default()
}

func (srv *Server) Start() {
	srv.engine.run()
}

func (srv *Server) SetupHandler(method string, route string, handler *gin.Handler) {
	srv.engine.addRoute(method, route, handler)
}
