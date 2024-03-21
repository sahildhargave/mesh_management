package main

import (
	"os"
	middleware "restruant-management/Middleware"
	"restruant-management/database"
	"restruant-management/routes"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	time.Now()

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	//routes.FoodRoutes(router)
	//routes.OrderRoutes(router)
	//routes.MenuRoutes(router)
	//routes.TableRoutes(router)
	//routes.OrderItemRoutes(router)
	//routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
