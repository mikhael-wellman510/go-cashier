package seeders

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type databaseConnection struct {
	db *gorm.DB
}

func NewSeeders(db *gorm.DB) *databaseConnection {

	return &databaseConnection{
		db: db,
	}
}

func (dc *databaseConnection) GenerateSeeders(ctx *gin.Context) {

	log.Println("tes")
}
