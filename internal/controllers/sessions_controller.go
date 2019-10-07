package controllers

import (
	"context"
	"time"
	"github.com/vasanthpandia/gojournal/internal/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionsController struct {
	Client *mongo.Client
	Collection string
	JwtKey []byte
}

func (sc *SessionsController) fetchUser(payload *LoginPayload) (*models.User, error) {
	filter := bson.D{{ "username", payload.Username }}
	var user models.User

	collection := sc.Client.Database("gojournal").Collection("users")

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (sc *SessionsController) Login(payload *LoginPayload) (string, error) {

	user, err := sc.fetchUser(payload)

	if err != nil {
		return "", err
	}

	if user.Authenticate(payload.Password) != nil {
		return "", err
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claim {
		Username: payload.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(sc.JwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
