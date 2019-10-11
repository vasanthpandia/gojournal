package middleware

import (
	"fmt"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/vasanthpandia/gojournal/internal/controllers"
	"github.com/vasanthpandia/gojournal/internal/models"
	"github.com/vasanthpandia/gojournal/internal/config"
)



func RequireAuth(cfg *config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.MustGet("Logger").(*zap.Logger)

		tokenstr := c.Request.Header.Get("X-Authentication")

		token, err := jwt.ParseWithClaims(tokenstr, &controllers.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return cfg.Token.Key, nil
		})

		if claims, ok := token.Claims.(*controllers.Claim); ok && token.Valid {
			logger.Info("Token", zap.String("username", claims.Username))

			filter := bson.D{{ "username", claims.Username }}
			var user models.User

			collection := cfg.DBConnection.Database.Collection("users")

			err := collection.FindOne(context.TODO(), filter).Decode(&user)
			if err != nil {
				c.JSON(400, gin.H{
					"error" : "Invalid Token",
				})
				c.Abort()
			}

			c.Set("CurrentUser", &user)
			c.Next()
		} else {
			fmt.Println(err)
			logger.Error("Token Error", zap.Error(err))
			c.JSON(403, "Token Expired")
			c.Abort()
		}
	}
}

func SetupLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		c.Set("Logger", logger)
		c.Next()
	}
}
