package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saadrupai/order-assignment/app/config"
	"github.com/saadrupai/order-assignment/app/models"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):] // Remove Bearer prefix
		claims := &models.Claims{}
		config := config.LoadConfig()
		jwtKey := []byte(config.SecretKey)
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("username", claims.UserEmail)
		c.Next()
	}
}
