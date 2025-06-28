package migrations

import (
	"log"
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	err := db.AutoMigrate(&entities.User{}, &entities.Store{}, &entities.Categories{}, &entities.Product{}, &entities.PaymentMethod{})

	if err != nil {
		return err
	}

	log.Println("Migration Succes")
	return nil
}
