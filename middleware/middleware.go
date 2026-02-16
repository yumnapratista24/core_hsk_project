package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format. Expected: Bearer <token>",
			})
			return
		}

		token := parts[1]
		expectedToken := os.Getenv("APIKey")

		if expectedToken == "" {
			log.Printf("WARNING: APIKey environment variable not set")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server configuration error",
			})
			return
		}

		if token != expectedToken {
			log.Printf("Authentication failed: invalid token from %s", c.ClientIP())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		c.Set("authenticated", true)
		log.Printf("Authentication successful from %s", c.ClientIP())
		c.Next()
	}
}
