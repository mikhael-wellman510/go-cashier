package migrations

import (
	"log"
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	err := db.AutoMigrate(&entities.Users{}, &entities.Store{}, &entities.Categories{})

	if err != nil {
		return err
	}

	log.Println("Migration Succes")
	return nil
}
