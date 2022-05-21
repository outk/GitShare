package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outk/gifshare/app/domain"
)

func main() {
	r := gin.Default()

	contents := r.Group("contents")
	{
		contents.GET("", func(ctx *gin.Context) {
			ctx.String(200, "contents")
		})
		contents.POST("", func(ctx *gin.Context) {
			var content domain.Content
			if err := ctx.ShouldBind(&content); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		})
	}

	r.Run()
}
