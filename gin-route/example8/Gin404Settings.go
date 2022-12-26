package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		gin设置找不到路由之后的404页面
	*/

	// 创建engine
	engine := gin.Default()
	// 处理get请求
	engine.GET("/user", func(context *gin.Context) {
		//
		name := context.DefaultQuery("name", "zhangsan")
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    name,
		})
	})
	// 设置找不到对应的路由后
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "404 not found",
		})
	})
	// 启动engine
	engine.Run(":8000")
}
