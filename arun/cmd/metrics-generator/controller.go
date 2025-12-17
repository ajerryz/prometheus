package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func RouteRegister(r *gin.Engine) {
	userG := r.Group("/user")
	{
		uc := new(UserController)
		userG.GET("/list", uc.UserList)
		userG.POST("/create", uc.UserList)
		userG.POST("/delete", uc.UserList)
		userG.POST("/update", uc.UserList)
		userG.GET("/user/:userId", uc.UserList)
	}
}

type UserController struct {
}

func (c UserController) UserList(ctx *gin.Context) {
	sleepMs := rand.Intn(2000)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	ctx.JSON(200, gin.H{
		"message": "User List",
	})
}
