package main

import (
	"./controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		pwd := c.PostForm("pwd")
		fmt.Print(pwd)
		result := controller.LoginHandler(username, pwd)
		if result {
			c.JSON(200, gin.H{
				"status": "OK",
				"msg": "login successfully",
			})
		} else {
			c.JSON(401, gin.H{
				"status": "Unauthorized",
				"msg": "login failed",
			})
		}
	})
	router.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		pwd := c.PostForm("pwd")
		result := controller.RegisterHandler(username, pwd)
		if result == "OK" {
			c.JSON(200, gin.H{
				"status": "OK",
				"msg": result,
			})
		} else {
			c.JSON(401, gin.H{
				"status": "Unauthorized",
				"msg": result,
			})
		}
	})
	router.Run(":8080")
}
