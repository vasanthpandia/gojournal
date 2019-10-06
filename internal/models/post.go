package models

import (
	"time"
)

type Post struct {
	ID string `json:"id" bson:"_id"`
	Date time.Time `json:"date" bson:"date"`
	Text string `json:"text" bson:"text"`
	Revision int `json:"-" bson:"revision"`
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	UpdatedAt time.Time `json:"-" bson:"UpdatedAt"`
}
