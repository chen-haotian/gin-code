package main

import "github.com/gin-gonic/gin"

func main() {
	/*
		两种创建Engine的方式
		第一种：使用gin.Default()
		第二种：使用gin.New()
	*/
	engine1 := gin.Default()
	engine1.Run(":8000")
	//engine2 := gin.New()
	//engine2.Run(":8000")
}
