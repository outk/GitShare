package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outk/gifshare/app/domain"
)

func GetGif(ctx *gin.Context) {
	ctx.String(200, "gif")
}

func PostGif(ctx *gin.Context) {
	var content domain.Content
	ctx.ShouldBind(&content)
	ctx.JSON(http.StatusOK, gin.H{
		"gif": "This is a gif.",
	})
}
