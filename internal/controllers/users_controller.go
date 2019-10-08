package controllers

import (
	"errors"
	"context"
	"github.com/vasanthpandia/gojournal/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type UsersController struct {
	Client *mongo.Client
	Collection string
}

type UserCreatePayload struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type UpdatePayload struct {
	UserCreatePayload
}

func (uc *UsersController) Create(payload *UserCreatePayload) (*models.User, error) {
	if payload.Password != payload.PasswordConfirmation {
		return nil, errors.New("Password Mismatch")
	}

	user := models.NewUser()
	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	user.Username = payload.Username
	user.Email = payload.Email

	user.BuildPassword(payload.Password)

	collection := uc.Client.Database("gojournal").Collection("users")

	_, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UsersController) Read(userId string) (*models.User, error) {
	filter := bson.D{{ "_id", userId }}
	var user models.User

	collection := uc.Client.Database("gojournal").Collection("users")

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
			return nil, err
	}

	return &user, nil
}
