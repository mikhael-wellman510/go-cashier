package repositories

import (
	"log"
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

type (
	StoreRepository interface {
		Create(entities.Store) (entities.Store, error)
		FindById(id string) (entities.Store, error)
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

func (sr *storeRepository) Create(store entities.Store) (entities.Store, error) {

	err := sr.db.Create(&store).Error
	log.Println("Err repo : ", err)

	return store, err
}

func (sr *storeRepository) FindById(id string) (entities.Store, error) {
	var store entities.Store

	err := sr.db.First(&store, "id=?", id).Error
	log.Println("Hasil err repo : ", err)
	log.Println("hasil store repo : ", store)
	return store, err
}
