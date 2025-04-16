package entities

import "time"

type Store struct {
	Base
	StoreName string `json:"storeName" gorm:"column:store_name;not null;unique"`
	Address   string `json:"address" gorm:"column:address;not null;unique"`
	OwnerName string `json:"ownerName" gorm:"column:owner_name;not null"`
}

type (
	StoreRequest struct {
		Id        string `json:"id"`
		StoreName string `json:"storeName" binding:"required"`
		Address   string `json:"address" binding:"required"`
		OwnerName string `json:"ownerName" binding:"required"`
	}

	StoreResponse struct {
		Id        string    `json:"id"`
		StoreName string    `json:"storeName"`
		Address   string    `json:"address"`
		OwnerName string    `json:"ownerName"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
