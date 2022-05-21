package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/outk/gifshare/app/domain"
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
