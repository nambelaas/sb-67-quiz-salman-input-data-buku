package structs

import (
	"time"
)

type Book struct {
	Id          int
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	ImageUrl    *string `json:"image_url"`
	ReleaseYear int     `json:"release_year" binding:"required,gte=1980,lte=2024"`
	Price       int     `json:"price" binding:"required,gt=0"`
	TotalPage   int     `json:"total_page" binding:"required,gt=0"`
	Thickness   string  `json:"thickness"`
	CategoryId  int     `json:"category_id" binding:"required"`
	CreatedAt   time.Time
	CreatedBy   string
	ModifiedAt  *time.Time
	ModifiedBy  *string
}
