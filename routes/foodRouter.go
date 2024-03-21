package routes

import (
	controller "restruant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	router.GET("/foods", controller.GetFoods())
	router.GET("/foods/:food_id", controller.GetFood())
	router.POST("/foods", controller.CreateFood())
	router.PATCH("/foods/:food_id", controller.UpdateFood())
}
