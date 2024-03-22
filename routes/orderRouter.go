package routes

import (
	controller "restruant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	router.GET("/orders", controller.GetOrders())
	router.GET("/orders/:order_id", controller.GetOrder())
	router.POST("/orders", controller.CreateOrder())
	router.PATCH("/orders/:order_id", controller.UpdateOrder())
}
