package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Route *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Route: gin.Default(),
	}
}

func (srv *Server) Start() {
	srv.Route.Run()
}
