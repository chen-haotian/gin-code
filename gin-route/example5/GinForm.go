package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		form表单参数
		使用context.PostForm()来接收form表单提交过来的参数
	*/

	// 创建engine
	engine := gin.Default()
	// 处理post请求
	engine.POST("/form", func(context *gin.Context) {
		UserName := context.PostForm("UserName")
		Password := context.PostForm("Password")
		fmt.Println(UserName, Password)
		// 响应请求
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    UserName + "登录成功!",
		})
	})
	// engine启动
	engine.Run(":8000")
}
