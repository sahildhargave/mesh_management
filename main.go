package main

import (
	"os"
	"time"
    "restruant-managment/database"
	"restruant-managment/routes"
	"restruant-managment/Middleware"
	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.foodCollection = database.OpenCollection(database.Client,"food")


func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}
	time.Now()

	router := gin.New()
	router.Use(gin.Logger())

	router.UserRouter(router)
	router.Use(Middleware.Authentication())
    
	
	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	
    
    router.Run(":"+ port)
}