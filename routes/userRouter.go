package routes

import (
	controller "restruant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/user", controller.GetUser())
	router.GET("/users/:user_id", controller.GetUser())
	router.POST("/users/signup", controller.SignUp())
	router.POST("/users/login", controller.Login())
}
