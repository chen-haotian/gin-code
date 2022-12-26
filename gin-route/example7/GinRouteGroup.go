package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		路由组
		相同路由接口的可以放在同一个路由组下
	*/

	// 创建默认的engine实例
	engine := gin.Default()
	v1 := engine.Group("v1")
	v2 := engine.Group("v2")

	{
		v1.GET("/i1", Info)
		v1.POST("/i2", Hello)
	}

	{
		v2.GET("/i1", Info)
		v2.POST("/i2", Hello)
	}
	// 启动engine
	engine.Run(":8000")
}

func Info(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "hello world",
	})
}
