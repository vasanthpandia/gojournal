package main

import(
	"fmt"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"github.com/vasanthpandia/gojournal/internal/config"
	"github.com/vasanthpandia/gojournal/internal/handlers"
	"github.com/vasanthpandia/gojournal/internal/server"
	"github.com/vasanthpandia/gojournal/internal/controllers"
)

func main() {
	cfg := config.InitDefaults()
	// fmt.Println(cfg.Mongo.Url)
	// fmt.Println(cfg.Mongo.Database)

	srv := server.NewServer()
	client, err := getMongoClient(cfg)

	if err != nil {
		fmt.Println(err)
	}

	route := srv.Route

	route.Use(setupLogger())

	route.Use(setupControllers(client))
	route.GET("/test", handlers.BasicHandler)
	route.POST("/users", handlers.CreateUser)
	srv.Start()
}

func setupControllers(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("UsersController", &controllers.UsersController{
			Client: client,
			Collection: "users",
		})
		c.Next()
	}
}

func getMongoClient(cfg *config.Config) (*mongo.Client, error){
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

func setupLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		c.Set("Logger", logger)
		c.Next()
	}
}
