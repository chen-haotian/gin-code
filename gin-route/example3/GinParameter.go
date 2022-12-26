package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		gin参数解析
		get请求解析的url：http://127.0.0.1:8000/student/zhangsan
	*/

	// 创建engine实例
	engine := gin.Default()
	// 处理get请求
	engine.GET("/student/:name", func(context *gin.Context) {
		name := context.Param("name")
		fmt.Println(name)
		// 响应请求
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "get success",
			"data":    "姓名=" + name,
		})
	})
	// 处理post请求
	engine.POST("/student/:name", func(context *gin.Context) {
		name := context.Param("name")
		fmt.Println(name)
		// 响应请求
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "post success",
			"data":    "姓名=" + name,
		})
	})
	// 启动engine
	engine.Run(":8000")
}
