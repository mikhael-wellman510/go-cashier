package repositories

import (
	"log"
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

type (
	CategoriesRepository interface {
		Create(categories *entities.Categories) error
		FindById(id string) (entities.Categories, error)
	}

	categoriesRepository struct {
		db *gorm.DB
	}
)

func NewCategoriesRepository(db *gorm.DB) CategoriesRepository {

	return &categoriesRepository{
		db: db,
	}
}

func (cr *categoriesRepository) Create(categories *entities.Categories) error {

	log.Println("arti cr : ", cr.db)
	return cr.db.Create(&categories).Error
}

func (cr *categoriesRepository) FindById(id string) (entities.Categories, error) {

	var categories entities.Categories

	err := cr.db.First(&categories, "id=?", id).Error

	return categories, err
}
