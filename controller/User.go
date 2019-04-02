package controller

import (
	"../model"
	"github.com/gin-gonic/gin"
)

type UserBody struct {
	Username string `json:"username" binding:"required"`
	Pwd string `json:"pwd" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var reqBody UserBody
	err := c.BindJSON(&reqBody)
	username := reqBody.Username
	pwd := reqBody.Pwd
	result := model.UserAuth(username, pwd)
	if err == nil && result {
		c.JSON(200, gin.H{
			"status": "OK",
			"msg": "login successfully",
		})
	} else if err == nil {
		c.JSON(401, gin.H{
			"status": "Unauthorized",
			"msg": "login failed",
		})
	} else {
		c.JSON(400, gin.H{
			"status": "Bad Request",
			"msg": "invalid params",
		})
	}
}

func RegisterHandler(c *gin.Context) {
	var reqBody UserBody
	err := c.BindJSON(&reqBody)
	username := reqBody.Username
	pwd := reqBody.Pwd
	result := model.UserRegister(username, pwd)
	if err == nil && result == "OK" {
		c.JSON(200, gin.H{
			"status": "OK",
			"msg": "register successfully",
		})
	} else if err == nil {
		c.JSON(403, gin.H{
			"status": "Unauthorized",
			"msg": result,
		})
	} else {
		c.JSON(400, gin.H{
			"status": "Bad Request",
			"msg": "invalid params",
		})
	}
}