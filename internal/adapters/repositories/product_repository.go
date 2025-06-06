package repositories

import (
	"log"
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Create(product *entities.Product) error
		Update(product *entities.Product) error
		FindById(id string) (*entities.Product, error)
		FindAllPagging(page int, limit int, search string) ([]entities.Product, error)
		CountWithFilterProduct(search string) (int64, error)
		FindProductByDate(date string) ([]entities.Product, error)
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
	return pr.db.Updates(product).Error
}

func (pr *productRepository) FindById(id string) (*entities.Product, error) {
	var product entities.Product

	err := pr.db.Preload("Categories").Preload("Store").First(&product, "id=?", id).Error

	return &product, err

}

func (pr *productRepository) FindAllPagging(page int, limit int, search string) ([]entities.Product, error) {
	var product []entities.Product
	offset := (page - 1) * limit
	query := pr.db.Model(&entities.Product{}).Preload("Categories").Preload("Store")

	if search != "" {
		query = query.Where("product_name LIKE ? or barcode LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query = query.Limit(limit).Offset(offset)

	if err := query.Find(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (pr *productRepository) CountWithFilterProduct(search string) (int64, error) {
	var count int64
	queryCount := pr.db.Model(&entities.Product{})

	if search != "" {
		queryCount = queryCount.Where("product_name LIKE ? or barcode LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := queryCount.Count(&count).Error
	log.Println("hasil nya adalah : ", count)
	return count, err
}

func (pr *productRepository) FindProductByDate(date string) ([]entities.Product, error) {

	var product []entities.Product

	err := pr.db.Where("DATE(created_at) = ?", date).Find(&product).Error

	return product, err
}
