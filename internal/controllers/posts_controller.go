package controllers

import (
	"context"
	"time"
	"github.com/vasanthpandia/gojournal/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type PostsController struct {
	Client *mongo.Client
	Collection *mongo.Collection
}

type PostCreatePayload struct {
	UserID string `json:"userId"`
	Date string `json:"date"`
	Text string `json:"text"`
}

type PostReadPayload struct {
	UserID string `json:"userId"`
	ID string `json:"id"`
}

type PostDeletePayload struct {
	UserID string `json:"userId"`
	ID string `json:"id"`
}

func (pc *PostsController) Create(payload *PostCreatePayload) (*models.Post, error) {
	post := models.NewPost()

	post.UserID = payload.UserID
	t, err := time.Parse(time.RFC3339, payload.Date)
	if err != nil {
		return nil, err
	}
	post.Date = t
	post.Text = payload.Text

	_, err = pc.Collection.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pc *PostsController) Read(payload *PostReadPayload) (*models.Post, error) {
	var post models.Post

	filter := bson.D{{"userId", payload.UserID}, {"_id", payload.ID}}

	err := pc.Collection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (pc *PostsController) Delete(payload *PostDeletePayload) error {
	filter := bson.D{{"userId", payload.UserID}, {"_id", payload.ID}}

	_, err := pc.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}