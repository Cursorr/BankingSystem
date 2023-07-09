package models

import (
	"Cursorr/BankingSystem/middleware"
	"Cursorr/BankingSystem/routes"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func (r *Router) SetupRoutes() {
	r.GET("/", routes.IndexRoute())
	r.GET("/validate", middleware.RequireAuth, routes.ValidateRoute())
	r.GET("/users", middleware.RequireAdmin, routes.UsersRoute())

	r.POST("/signup", routes.SignupRoute())
	r.POST("/login", routes.LoginRoute())
}

func NewRouter() *Router {
	router := gin.Default()
	return &Router{router}
}
