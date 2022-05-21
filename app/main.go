package main

import (
	"gifshare/app/controller"

	"github.com/gin-gonic/gin"
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
