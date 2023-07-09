package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersRoute() gin.HandlerFunc {
	return Users
}

func Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "Users API",
	})
}
