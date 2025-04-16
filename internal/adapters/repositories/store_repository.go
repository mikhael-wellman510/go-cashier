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
		Update(id string, update entities.Store) (entities.Store, error)
		Deleted(id string) error
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

func (sr *storeRepository) Update(id string, update entities.Store) (entities.Store, error) {

	var store entities.Store

	if err := sr.db.First(&store, "id=?", id).Error; err != nil {
		return store, err
	}

	if err := sr.db.Model(&store).Updates(update).Error; err != nil {
		return store, err

	}

	return store, nil
}

func (sr *storeRepository) Deleted(id string) error {

	var store entities.Store

	// Kalau pake uuid harus pakai id=? , karena dia mengira itu integer
	return sr.db.Delete(&store, "id=?", id).Error
}
