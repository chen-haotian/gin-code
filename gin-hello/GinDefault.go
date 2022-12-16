package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建默认engine
	engine := gin.Default()
	// engine调用Handle方法处理GET请求
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		// 响应请求
		// gin.H 返回的是一个map类型的数据结构
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    "hello gin",
		})
	})
	// engine启动
	engine.Run(":8000")
}
