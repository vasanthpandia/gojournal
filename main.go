package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vasanthpandia/gojournal/internal/config"
	"github.com/vasanthpandia/gojournal/internal/controllers"
	"github.com/vasanthpandia/gojournal/internal/handlers"
	"github.com/vasanthpandia/gojournal/internal/server"
	"github.com/vasanthpandia/gojournal/internal/middleware"
)

func main() {
	cfg := config.InitDefaults()

	srv := server.NewServer()
	client, err := getDBClient(cfg)

	if err != nil {
		fmt.Println(err)
	}

	route := srv.Route

	route.Use(middleware.SetupLogger())

	route.Use(setupControllers(client))
	route.GET("/test", handlers.BasicHandler)
	route.POST("/users", handlers.CreateUser)
	route.POST("/login", handlers.Login)
	route.Use(middleware.RequireAuth(client))
	route.GET("/users/:userId", handlers.GetUser)
	route.POST("/posts", handlers.CreatePost)
	route.GET("/posts/:postId", handlers.ReadPost)
	route.DELETE("/posts/:postId", handlers.DeletePost)
	srv.Start()
}

func setupControllers(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("UsersController", &controllers.UsersController{
			Client:     client,
			Collection: "users",
		})
		c.Set("SessionsController", &controllers.SessionsController{
			Client:     client,
			Collection: "users",
			JwtKey:     []byte("DEFAULTKEY"),
		})
		c.Set("PostsController", &controllers.PostsController{
			Client: client,
			Collection: "posts",
		})
		c.Next()
	}
}

func getDBClient(cfg *config.Config) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(cfg.Mongo.Url)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("Unable to Initialize Mongo Client")
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}
