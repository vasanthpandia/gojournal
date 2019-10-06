package main

import(
	"fmt"

	"github.com/vasanthpandia/gojournal/internal/config"
	"github.com/vasanthpandia/gojournal/internal/handlers"
	"github.com/vasanthpandia/gojournal/internal/server"
)

func main() {
	cfg := config.InitDefaults()
	fmt.Println(cfg.Mongo.Url)
	fmt.Println(cfg.Mongo.Database)
	srv := server.NewServer()
	route := *srv.Route

	route.GET("/test", handlers.BasicHandler)
	srv.Start()
}
