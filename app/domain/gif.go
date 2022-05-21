package domain

type Content struct {
	Name Name `json:"name" binding:"required"`
	Gif  Gif  `json:"gif" binding:"required"`
}

type Name string

type Gif struct{}
