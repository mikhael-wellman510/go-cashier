package entities

import "time"

type Categories struct {
	Base
	CategoryName string `json:"category_name" gorm:"column:category_name;not null;unique"`
	Description  string `json:"description" gorm:"column:description;not null;unique"`
	ImageUrl     string `json:"image_url" gorm:"column:image_url"`
}

type (
	CategoryRequest struct {
		Id           string `form:"id"`
		CategoryName string `form:"categories_name"`
		Description  string `form:"description"`
		ImageUrl     string
	}

	CategoryResponse struct {
		Id           string    `json:"id"`
		CategoryName string    `json:"category_name"`
		Description  string    `json:"description"`
		ImageUrl     string    `json:"image_url"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
