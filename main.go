package main

import (
	"github.com/gin-gonic/gin"

	"github.com/vasanthpandia/gojournal/internal/config"
	"github.com/vasanthpandia/gojournal/internal/controllers"
	"github.com/vasanthpandia/gojournal/internal/handlers"
	"github.com/vasanthpandia/gojournal/internal/middleware"
	"github.com/vasanthpandia/gojournal/internal/server"
)

func main() {
	cfg := config.GetServerConfig("development")

	srv := server.NewServer()
	route := srv.Route

	route.Use(middleware.SetupLogger())
	route.Use(setupControllers(cfg))

	route.GET("/test", handlers.BasicHandler)
	route.POST("/users", handlers.CreateUser)
	route.POST("/login", handlers.Login)

	route.Use(middleware.RequireAuth(cfg))
	route.GET("/users/:userId", handlers.GetUser)
	route.POST("/posts", handlers.CreatePost)
	route.GET("/posts/:postId", handlers.ReadPost)
	route.DELETE("/posts/:postId", handlers.DeletePost)

	srv.Start()
}

func setupControllers(cfg *config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("UsersController", &controllers.UsersController{
			Client:     cfg.DBConnection.Client,
			Collection: cfg.DBConnection.Database.Collection("users"),
		})
		c.Set("SessionsController", &controllers.SessionsController{
			Client:     cfg.DBConnection.Client,
			Collection: cfg.DBConnection.Database.Collection("users"),
			JwtKey:     cfg.Token.Key,
			Validity:   cfg.Token.Validity,
		})
		c.Set("PostsController", &controllers.PostsController{
			Client:     cfg.DBConnection.Client,
			Collection: cfg.DBConnection.Database.Collection("posts"),
		})
		c.Next()
	}
}
