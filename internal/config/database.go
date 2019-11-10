package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Client *mongo.Client
	Database *mongo.Database
}

func GetMongoConnection(config *MongoConfig) (*MongoConnection, error) {
	clientOptions := options.Client().ApplyURI(config.Url)

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

	database := client.Database(config.Database)

	fmt.Println("Connected to MongoDB!")

	connection := &MongoConnection {
		Client: client,
		Database: database,
	}

	return connection, nil
}
