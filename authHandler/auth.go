package authhandler

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"../auth" 


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader { 
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}

		if !auth.ValidateToken(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Next() 
	}
}