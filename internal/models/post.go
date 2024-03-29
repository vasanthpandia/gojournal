package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    string    `json:"userId" bson:"userId"`
	Date      time.Time `json:"date" bson:"date"`
	Title     string    `json:"title" bson:"title"`
	Text      string    `json:"text" bson:"text"`
	Revision  int       `json:"-" bson:"revision"`
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	UpdatedAt time.Time `json:"-" bson:"UpdatedAt"`
}

func NewPost() *Post {
	now := time.Now()
	return &Post{
		ID:        uuid.New().String(),
		Revision:  1,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
