package middleware

import (
	"Cursorr/BankingSystem/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := getUserFromToken(tokenString)
	if err != nil || user == (database.User{}) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	c.Next()
}
