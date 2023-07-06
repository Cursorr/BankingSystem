package models

import (
	"Cursorr/BankingSystem/routes"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func (r *Router) SetupRoutes() {
	r.GET("/", routes.IndexRoute())
	r.POST("/signup", routes.SignupRoute())
}

func NewRouter() *Router {
	router := gin.Default()
	return &Router{router}
}
