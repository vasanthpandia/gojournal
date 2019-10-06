package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
	// // "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/options"
)

type BaseController struct {
	client *mongo.Client
}

func BasicController() string {
	return "Controller Working!"
}
