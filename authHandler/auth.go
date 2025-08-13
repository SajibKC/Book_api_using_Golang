package authhandler

import (
	"encoding/base64"
	"example/Book_api_using_Golang/dataHandler"
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		log.Println("DEBUG: Raw Authorization header =", authHeader)

		token := strings.TrimPrefix(authHeader, "Basic ")
		if token == authHeader { 
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}
		decodedBytes, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid base64 token"})
			return
		}
		decodedStr := string(decodedBytes)
		parts := strings.SplitN(decodedStr, ":", 2)
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}
		username, password := parts[0], parts[1]
		var flag bool
		flag = false
		for _, user := range datahandler.Users {
			if user.Username == username && user.Password == password{
				flag = true
				break
			}
		}
		if !flag {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}


		c.Next() 
	}
}