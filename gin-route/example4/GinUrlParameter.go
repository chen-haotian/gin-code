package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		获取URL上的参数
		例如：http://127.0.0.1:8000/student?name=liri1 这个name就是URL上的参数
	*/
	// 创建engine实例
	engine := gin.Default()
	// get请求
	engine.GET("/student", func(context *gin.Context) {
		// DefaultQuery()
		name := context.DefaultQuery("name", "张三")
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "get success",
			"data":    name + "登录成功!",
		})
	})
	// post请求
	engine.POST("/student", func(context *gin.Context) {
		// Query()
		name := context.Query("name")
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "post success",
			"data":    name + "登录成功!",
		})
	})
	engine.Run(":8000")
}
