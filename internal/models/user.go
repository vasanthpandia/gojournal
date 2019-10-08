package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

type User struct {
	ID string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	Email string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	HashedPassword string `json:"-" bson:"hashedPassword"`
	DateOfBirth time.Time `json:"dateOfBirth" bson:"dateOfBirth"`
	Revision int `json:"-" bson:"revision"`
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	UpdatedAt time.Time `json:"-" bson:"UpdatedAt"`
	posts []Post
}

func NewUser() *User {
	now := time.Now()
	return &User {
		ID: uuid.New().String(),
		Revision: 1,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (user *User)Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
}

func (user *User)BuildPassword(password string) {
	user.HashedPassword = hashAndSalt([]byte(password))
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
			return "NoPasswordGenerated"
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
