package routes

import (
	controller "restruant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine) {
	router.GET("/menus", controller.GetMenus())
	router.GET("/menus/:menu_id", controller.GetMenu())
	router.POST("/menus", controller.CreateMenu())
	router.PATCH("/menus/:menu_id", controller.UpdateMenu())
}
