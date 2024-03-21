package routes

import "github.com/gin-gonic/gin"

func UserRouter(router *gin.Engine) {
	router.GET("/user", userController.GetUser())
}