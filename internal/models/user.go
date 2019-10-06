package models

import (
	"time"
)

type User struct {
	ID string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	Email string `json:"email" bson:"email"`
	HashedPassword string `json:"-" bson:"hashedPassword"`
	DateOfBirth time.Time `json:"dateOfBirth" bson:"dateOfBirth"`
}
