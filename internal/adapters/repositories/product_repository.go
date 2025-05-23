package repositories

import (
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Create(product *entities.Product) error
		Update(product *entities.Product) error
		FindById(id string) (entities.Product, error)
	}

	productRepository struct {
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {

	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) Create(product *entities.Product) error {

	return pr.db.Create(product).Error
}

func (pr *productRepository) Update(product *entities.Product) error {
	return pr.db.Save(product).Error
}

func (pr *productRepository) FindById(id string) (entities.Product, error) {
	var product entities.Product

	err := pr.db.Preload("Categories").Preload("Store").First(&product, "id=?", id).Error

	return product, err

}
