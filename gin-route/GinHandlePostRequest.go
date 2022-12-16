package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		使用Handle()方法来处理http请求
					第一个参数：表示处理的http请求的类型
					第二个参数：表示解析的接口路径URI
					第三个参数：表示处理对应的请求的函数
	*/
	// 创建默认engine
	engine := gin.Default()
	// Handle方法处理get请求
	engine.Handle("POST", "/hello", func(context *gin.Context) {
		// PostForm()可以解析POST请求Body携带的参数
		name := context.PostForm("name")
		password := context.PostForm("password")
		if name == "zhangsan" && password == "123456" {
			// 响应请求
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "success",
				"data":    name + "登录成功!",
			})
		}
	})
	// 启动engine
	engine.Run(":8000")
}
