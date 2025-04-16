package usecases

import (
	"log"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/entities"
)

type (
	StoreService interface {
		CreateStore(storeReq entities.StoreRequest) (entities.StoreResponse, error)
		FindStoreById(id string) (entities.StoreResponse, error)
	}

	storeService struct {
		StoreRepository repositories.StoreRepository
	}
)

// Constructor
func NewStoreService(storeRepository repositories.StoreRepository) StoreService {
	return &storeService{
		StoreRepository: storeRepository,
	}
}

func (ss *storeService) CreateStore(storeReq entities.StoreRequest) (entities.StoreResponse, error) {
	store := entities.Store{
		StoreName: storeReq.StoreName,
		Address:   storeReq.Address,
		OwnerName: storeReq.OwnerName,
	}

	res, err := ss.StoreRepository.Create(store)
	log.Println("Err useCase : ", err)
	if err != nil {
		return entities.StoreResponse{}, err
	}

	return entities.StoreResponse{
		Id:        res.ID,
		StoreName: res.StoreName,
		Address:   res.Address,
		OwnerName: res.OwnerName,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (ss *storeService) FindStoreById(id string) (entities.StoreResponse, error) {
	res, err := ss.StoreRepository.FindById(id)

	if err != nil {
		return entities.StoreResponse{}, err
	}

	return entities.StoreResponse{
		Id:        res.ID,
		StoreName: res.StoreName,
		Address:   res.Address,
		OwnerName: res.OwnerName,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, err
}
