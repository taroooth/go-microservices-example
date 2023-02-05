package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/taroooth/go-microservices-example/services/gateway/controller"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Failed to load env file: %v", err)
		log.Fatalln(err)
	}
	router := gin.Default()
	router.GET("/books", controller.GetBooks)
	router.POST("/books", controller.CreateBook)
	router.Run()
}
