package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	/*
		当个文件上传
		可以使用context.FormFile()来接收post提交的参数
	*/

	// 创建engine实例
	engine := gin.Default()
	// 限制上传文件大小为
	engine.MaxMultipartMemory = 8 << 20 // 8Mb
	// 处理post请求提交的数据
	engine.POST("/upload", func(context *gin.Context) {
		// 用于接收file
		file, err := context.FormFile("file")
		// 用于拼接文件上传的路径
		dst := path.Join("./gin-route/file/" + file.Filename)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "failed",
			})
			return
		}
		err = context.SaveUploadedFile(file, dst)
		// 上传成功响应
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	})
	// 启动engine
	engine.Run(":8000")
}
