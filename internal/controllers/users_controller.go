package controllers

import (
	"errors"
	"github.com/vasanthpandia/gojournal/internal/models"
)

type UsersController struct {
	BaseController
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

func (uc *UsersController) create(payload UserCreatePayload) (*models.User, error) {
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
