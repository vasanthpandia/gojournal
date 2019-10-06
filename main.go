package main

import(
	"github.com/vasanthpandia/gojournal/internal/handlers"
	"github.com/vasanthpandia/gojournal/internal/server"
)

func main() {
	srv := server.NewServer()
	route := *srv.Route

	route.GET("/test", handlers.BasicHandler)
	srv.Start()
}
