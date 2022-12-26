package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	/*
		多文件上传
	*/
	// 创建engine实例
	engine := gin.Default()
	// 限制表单上传大小 8MB，默认为32MB
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/upload", func(context *gin.Context) {
		multipartForm, err := context.MultipartForm()
		if err != nil {
			// 文件读取失败的响应
			context.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "filed",
			})
			return
		}
		// 获取所有图片
		files := multipartForm.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 按照指定路径存储图片
			dst := path.Join("./gin-route/file/" + file.Filename)
			if err := context.SaveUploadedFile(file, dst); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"message": "filed",
				})
			}
		}
		// 上传成功响应
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	})
	// 启动engine
	engine.Run(":8000")
}
