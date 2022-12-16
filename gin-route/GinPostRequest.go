package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		post请求
	*/
	// 创建engine实例
	engine := gin.Default()
	// 处理post请求
	engine.POST("/hello", func(context *gin.Context) {
		// PostForm方法是用于解析post请求form-data里面的参数，form表单
		UserName := context.PostForm("UserName")
		Password := context.PostForm("Password")
		fmt.Println(UserName)
		fmt.Println(Password)
		// 响应处理
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    UserName + "登录成功",
		})
	})
	// 启动engine
	engine.Run(":8000")
}
