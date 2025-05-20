package entities

import "github.com/shopspring/decimal"

type (
	Product struct {
		Base
		Stock        int             `json:"stock" gorm:"column:category;not null"`
		Barcode      string          `json:"barcode" gorm:"column:barcode;not null"`
		ProductName  string          `json:"product_name" gorm:"column:product_name;not null"`
		Description  string          `json:"description" gorm:"column:description;not null"`
		Price        decimal.Decimal `json:"price" gorm:"column:price;not null"`
		CategoriesId string          `json:"categories_id" gorm:"column:categories_id"`
		StoreId      string          `json:"store_id" gorm:"column:store_id"`
		Categories   Categories      `gorm:"foreignKey:CategoriesId"`
		Store        Store           `gorm:"foreignKey:StoreId"`
	}

	ProductRequest struct {
		Id           string          `json:"id"`
		Stock        int             `json:"stock"`
		Barcode      string          `json:"barcode"`
		ProductName  string          `json:"product_name"`
		Description  string          `json:"description"`
		Price        decimal.Decimal `json:"price"`
		CategoriesId string          `json:"categories_id"`
		StoreId      string          `json:"store_id"`
	}

	ProductResponse struct {
		Id               string           `json:"id"`
		Stock            int              `json:"stock"`
		Barcode          string           `json:"barcode"`
		ProductName      string           `json:"product_name"`
		Description      string           `json:"description"`
		Price            decimal.Decimal  `json:"price"`
		CategoryResponse CategoryResponse `json:"category"`
		StoreResponse    StoreResponse    `json:"store"`
	}
)
