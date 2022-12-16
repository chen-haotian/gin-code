package main

import (
	"fmt"
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

	// Handle()通用处理请求处理get请求
	// 创建默认engine
	engine := gin.Default()
	// Handle方法处理get请求
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		// 获取请求的uri路径【解析请求的uri路径】
		path := context.FullPath()
		fmt.Println(path) // 输出结果：/hello
		// 获取到请求拼接在uri上的参数，例如http://127.0.0.1:8080?name=zhangsan
		// 获取的参数不为nil就返回获取到的数据，否则就返回，参数二设置的defaultValue()的默认值
		query := context.DefaultQuery("name", "zhangsan")
		// 响应请求
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    "hello " + query,
		})
	})
	// 启动engine
	engine.Run(":8000")
}
