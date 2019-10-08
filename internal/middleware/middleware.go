package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/dgrijalva/jwt-go"

	"github.com/vasanthpandia/gojournal/internal/controllers"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		JwtKey := []byte("DEFAULTKEY")

		logger := c.MustGet("Logger").(*zap.Logger)

		tokenstr := c.Request.Header.Get("X-Authentication")

		token, err := jwt.ParseWithClaims(tokenstr, &controllers.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if claims, ok := token.Claims.(*controllers.Claim); ok && token.Valid {
			logger.Info("Token", zap.String("username", claims.Username))
			fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
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
