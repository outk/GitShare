package usecase

import (
	"gifshare/app/domain"

	"github.com/gin-gonic/gin"
)

func GetGif(ctx *gin.Context) (domain.Content, error) {
	return domain.Content{
		Name: "gif name",
		Gif:  domain.Gif{},
	}, nil
}

func CreateGif(ctx *gin.Context, content domain.Content) error {
	return nil
}
