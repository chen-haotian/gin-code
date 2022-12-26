package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		get请求示例
	*/
	// 获取engine实例
	engine := gin.Default()
	// 处理get请求
	// 解析url为http://127.0.0.1:8000/student?name=zhangsan&age=20
	engine.GET("/student", func(context *gin.Context) {
		// Query()方法直接解析获取get请求携带的参数
		name := context.Query("name")
		age := context.Query("age")
		// 响应请求
		if name == "zhangsan" && age == "20" {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "success",
				"data":    name + "成功登录!",
			})
		}
	})
	// 启动engine
	engine.Run(":8000")
}
