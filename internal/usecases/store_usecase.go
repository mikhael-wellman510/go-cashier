package usecases

import (
	"log"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/utils"
)

type (
	StoreService interface {
		CreateStore(storeReq entities.StoreRequest) (entities.StoreResponse, error)
		FindStoreById(id string) (entities.StoreResponse, error)
		UpdatedStore(storeReq entities.StoreRequest) (entities.StoreResponse, error)
		DeletedStore(id string) (bool, error)
		FilterAndPagginStore(page int, limit int, storeName string, ownerName string) (utils.PaginationResponse, error)
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

	// res, err := ss.StoreRepository.Create(&store)
	log.Println("store awal : ", store)
	if err := ss.StoreRepository.Create(&store); err != nil {
		log.Println("Error nya : ", err)
		return entities.StoreResponse{}, err
	}

	log.Println("Log store setelah create : ", store)
	return entities.StoreResponse{
		Id:        store.ID,
		StoreName: store.StoreName,
		Address:   store.Address,
		OwnerName: store.OwnerName,
		CreatedAt: store.CreatedAt,
		UpdatedAt: store.UpdatedAt,
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

func (ss *storeService) UpdatedStore(storeReq entities.StoreRequest) (entities.StoreResponse, error) {

	// Todo find by id dulu

	res, err := ss.StoreRepository.FindById(storeReq.Id)

	log.Println("Hasil res: ", res)
	if err != nil {
		log.Println("Hasil err : ", err)
		return entities.StoreResponse{}, err

	}

	res.StoreName = storeReq.StoreName
	res.Address = storeReq.Address
	res.OwnerName = storeReq.OwnerName

	if err := ss.StoreRepository.Update(&res); err != nil {

		return entities.StoreResponse{}, err
	}

	log.Println("Hasil updated : ", res)
	return entities.StoreResponse{
		Id:        res.ID,
		StoreName: res.StoreName,
		Address:   res.Address,
		OwnerName: res.OwnerName,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (ss *storeService) DeletedStore(id string) (bool, error) {
	_, err := ss.FindStoreById(id)

	if err != nil {
		return false, err
	}

	return true, ss.StoreRepository.Deleted(id)

}

func (ss *storeService) FilterAndPagginStore(page int, limit int, storeName string, ownerName string) (utils.PaginationResponse, error) {

	res, err := ss.StoreRepository.FindAllPagging(page, limit, storeName, ownerName)

	if err != nil {
		log.Println("Error in useCase : ", err.Error())
	}

	result, err := ss.StoreRepository.CountStoresWithFilter(storeName, ownerName)

	if err != nil {
		log.Println("Error count : ", result)
	}
	resPagging := utils.PaginationDto(res, int(result), page, limit)

	return resPagging, nil
}
