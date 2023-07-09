package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateRoute() gin.HandlerFunc {
	return Validate
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged IN",
	})
}
