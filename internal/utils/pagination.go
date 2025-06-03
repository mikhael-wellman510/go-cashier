package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PaginationResponse struct {
	Data         any  `json:"data"`          // Todo -> Data yang diterima
	TotalRecords int  `json:"total_records"` // Todo -> total jumlah data
	TotalPages   int  `json:"total_pages"`   // Todo -> Total halaman
	CurrentPages int  `json:"current_pages"` // Todo -> Halaman saat ini
	PerPage      int  `json:"per_page"`      // todo -> jumlah item per halaman
	HasNext      bool `json:"has_next"`      // todo -> apakah halaman berikut nya ada
	HasPrevious  bool `json:"has_previous"`  // todo -> apakah halaman sebelum nya ada
}

func PaginationDto(data any, totalRecords int, page int, limit int) PaginationResponse {

	// 20 total + 10 limit -1 = 29 / 10 = 2 pages
	// 21 total + 10 limit -1 = 30 / 10 = 3 pages
	totalPages := (totalRecords + limit - 1) / limit

	var hasNext bool = page < totalPages
	var hasPrevious bool = page > 1

	// Buat next page
	// Buat previous page
	return PaginationResponse{
		Data:         data,
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		CurrentPages: page,
		PerPage:      limit,
		HasNext:      hasNext,
		HasPrevious:  hasPrevious,
	}
}

func GetPagination(ctx *gin.Context) Pagination {

	// jika kosong , maka dia default nya 1
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	if err != nil || page == 0 {

		page = 1
	}

	// Jika kosong , maka dia default nya 10
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if err != nil {
		log.Println("err limit : ", err)
		limit = 10
	}

	return Pagination{
		Page:  page,
		Limit: limit,
	}
}
