package routes

import (
	controller "restruant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/user", controller.GetUser())
}
