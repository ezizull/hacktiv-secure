// Package middlewares contains the middlewares for the restapi api
package middlewares

import (
	"fmt"
	"net/http"

	"secure/challenge-2/utils/lists"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// AuthJWTMiddleware is a function that validates the jwt token
func AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		viper.SetConfigFile("config.json")
		if err := viper.ReadInConfig(); err != nil {
			_ = fmt.Errorf("fatal error in config file: %s", err.Error())
		}

		JWTAccessSecure := viper.GetString("Secure.JWTAccessSecure")
		tokenString := c.GetHeader("Authorization")
		signature := []byte(JWTAccessSecure)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return signature, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}

		fmt.Println("check ", c.Keys, claims["role"])
		fmt.Println("check ", claims)

		c.Keys["Role"] = claims["role"]
		c.Keys["UserID"] = claims["user_id"]

		c.Next()
	}
}

// AuthRoleMiddleware is a function that validates the role of user
func AuthRoleMiddleware(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
			c.Abort()
			return
		}

		roleVal := c.Keys["Role"].(string)
		if roleVal == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
			c.Abort()
			return
		}

		if !lists.Contains(validRoles, roleVal) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized for this path"})
			c.Abort()
			return
		}

		c.Next()
	}
}
