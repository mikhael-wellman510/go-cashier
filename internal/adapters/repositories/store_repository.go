package repositories

import (
	"log"
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

type (
	StoreRepository interface {
		Create(store *entities.Store) error
		FindById(id string) (entities.Store, error)
		Update(store *entities.Store) error
		Deleted(id string) error
		FindAllPagging(page int, limit int, storeName string, ownerName string) ([]entities.Store, error)
	}

	storeRepository struct {
		db *gorm.DB
	}
)

func NewStoreRepository(db *gorm.DB) StoreRepository {

	return &storeRepository{
		db: db,
	}
}

func (sr *storeRepository) Create(store *entities.Store) error {

	return sr.db.Create(&store).Error
}

func (sr *storeRepository) FindById(id string) (entities.Store, error) {
	var store entities.Store

	err := sr.db.First(&store, "id=?", id).Error

	return store, err
}

func (sr *storeRepository) Update(store *entities.Store) error {

	return sr.db.Save(store).Error
}

func (sr *storeRepository) Deleted(id string) error {

	var store entities.Store

	// Kalau pake uuid harus pakai id=? , karena dia mengira itu integer
	return sr.db.Delete(&store, "id=?", id).Error
}

func (sr *storeRepository) FindAllPagging(page int, limit int, storeName string, ownerName string) ([]entities.Store, error) {

	var stores []entities.Store

	query := sr.db.Model(&entities.Store{})

	if storeName != "" {
		query = query.Where("store_name LIKE ?", "%"+storeName+"%")
	}

	if ownerName != "" {
		query = query.Where("owner_name LIKE ?", "%"+ownerName+"%")
	}

	if err := query.Find(&stores).Error; err != nil {
		return nil, err
	}

	log.Println("Hasil stores : ->> ", stores)
	return []entities.Store{}, nil
}
