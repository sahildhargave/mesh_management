package main

import (
	"os"
)


func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}

	router := gin.New()

	router.User(gin.L)
}