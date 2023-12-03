package routes

import (
	controller "github.com/Suman196pokhrel/go-jwt-auth/controllers"
	"github.com/Suman196pokhrel/go-jwt-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.Use(middlewares.Authenticate)
	routes.GET("/users", controller.GetUsers)
	routes.GET("user/:user_id", controller.GetUser)
}
