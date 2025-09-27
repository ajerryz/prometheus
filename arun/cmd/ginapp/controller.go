package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) Hello(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		name = "world"
	}
	msg := fmt.Sprintf("hello %s", name)
	ctx.JSON(200, gin.H{
		"message": msg,
	})
}

func (c *HelloController) HelloPath(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

func (c *HelloController) Error400(ctx *gin.Context) {
	n := rand.Int64N(20)
	time.Sleep(time.Duration(n) * time.Second)
	ctx.JSON(400, gin.H{
		"message": "Error400",
	})
}

func (c *HelloController) Error500(ctx *gin.Context) {
	n := rand.Int64N(2000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	ctx.JSON(500, gin.H{
		"message": "Error500",
	})
}

type UserController struct {
}
