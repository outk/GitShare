package main

import (
	"github.com/gin-gonic/gin"
	"github.com/outk/gifshare/app/controller"
)

func main() {
	r := gin.Default()

	contents := r.Group("contents")
	{
		gif := contents.Group("gif")
		{
			gif.GET("", controller.GetGif)
			gif.POST("", controller.PostGif)
		}
	}

	r.Run()
}
