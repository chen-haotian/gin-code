package main

import "github.com/gin-gonic/gin"

func main() {
	/*
		两种创建Engine的方式
		第一种：使用gin.Default()
		第二种：使用gin.New()
	*/
	engine := gin.New()
	engine.Run(":8000")
}
