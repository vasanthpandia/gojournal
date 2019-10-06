package controllers

import (
	"errors"
	"github.com/vasanthpandia/gojournal/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// type UsersControllerInterface interface {
// 	Create(payload *UserCreatePayload) (*models.User, error)
// }

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

	user := &models.User {
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Username: payload.Username,
		Email: payload.Email,
	}

	user.BuildPassword(payload.Password)

	return user, nil
}
