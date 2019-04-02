package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/login", controller.LoginHandler)
	router.POST("/register", controller.RegisterHandler)
	router.Run(":8080")
}
