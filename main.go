package main

import(
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vasanthpandia/gojournal/internal/config"
	"github.com/vasanthpandia/gojournal/internal/handlers"
	"github.com/vasanthpandia/gojournal/internal/server"
)

func main() {
	cfg := config.InitDefaults()
	// fmt.Println(cfg.Mongo.Url)
	// fmt.Println(cfg.Mongo.Database)

	srv := server.NewServer()
	client := getMongoClient(cfg.Mongo)
	route := *srv.Route
	route.Use(setupControllers(client))
	route.GET("/test", handlers.BasicHandler)
	srv.Start()
}

func setupControllers(client *Mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		base := &BaseController {
			client: client,
		}

		c.Set("UsersController", base)
	}
}

func getMongoClient(mongo *config.Config.Mongo) (*Mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongo.Url)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		fmt.Println("Unable to Initialize Mongo Client")
	}

	return client
}
