package main

import(
	"fmt"
	"github.com/vasanthpandia/gojournal/internal/handlers"
	"github.com/vasanthpandia/gojournal/internal/server"
)

func main() {
	srv := server.GetDefaultServer()

	srv.SetupHandler("GET", "/test", handlers.BasicHandler)
	srv.Start()
}

// Add Logic for Building Config

// Initialize Server

// Add Controllers

// Add Routes
