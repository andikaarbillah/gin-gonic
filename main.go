package main

import (
	"apigin/controller"
	"apigin/middleware"
	"apigin/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func SetupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	SetupLogOutput()

	server := gin.Default()
	server.Use(gin.Recovery(), middleware.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(200, gin.H{"message": "Video Input is Succes"})
		}
	})

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": 200,
		})
	})

	server.Run()
}
