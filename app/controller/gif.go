package controller

import (
	"gifshare/app/domain"
	"gifshare/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getGifResponse struct {
	Name string `json:"name"`
}

type postGifRequest struct {
	Name string `json:"name" binding:"required"`
}

func GetGif(ctx *gin.Context) {
	content, err := usecase.GetGif(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	ctx.JSON(http.StatusOK, getGifResponse{
		Name: string(content.Name),
	})
	return
}

func PostGif(ctx *gin.Context) {
	var request postGifRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.String(http.StatusBadRequest, "bad request")
		return
	}

	if err := usecase.CreateGif(ctx, domain.Content{
		Name: domain.Name(request.Name),
	}); err != nil {
		ctx.String(http.StatusInternalServerError, "internal server error")
	}
	return
}
