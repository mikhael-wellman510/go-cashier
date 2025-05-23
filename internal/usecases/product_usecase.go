package usecases

import (
	"log"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/entities"
)

type (
	ProductService interface {
		CreateProduct(productReq entities.ProductRequest) (entities.ProductResponse, error)
		UpdateProduct(productReq entities.ProductRequest) (entities.ProductResponse, error)
	}

	// Class
	productService struct {
		ProductRepository repositories.ProductRepository
		StoreService      StoreService
		CategoryService   CategoryService
	}
)

// Kenapa dia return nya interface , karena ia ingin mengembalikan interface( method)
func NewProductService(productRepository repositories.ProductRepository, storeService StoreService, categoryService CategoryService) ProductService {
	return &productService{
		ProductRepository: productRepository,
		StoreService:      storeService,
		CategoryService:   categoryService,
	}
}

func (ps *productService) CreateProduct(productReq entities.ProductRequest) (entities.ProductResponse, error) {

	// find -> store id
	storeRes, err := ps.StoreService.FindById(productReq.StoreId)

	if err != nil {
		return entities.ProductResponse{}, err
	}

	categoryRes, err := ps.CategoryService.FindById(productReq.CategoriesId)

	if err != nil {
		return entities.ProductResponse{}, err
	}

	product := entities.Product{
		Stock:        productReq.Stock,
		Barcode:      productReq.Barcode,
		ProductName:  productReq.ProductName,
		Description:  productReq.Description,
		Price:        productReq.Price,
		CategoriesId: categoryRes.ID,
		StoreId:      storeRes.ID,
	}

	if err := ps.ProductRepository.Create(&product); err != nil {
		return entities.ProductResponse{}, err
	}
	log.Println("Hasil product : ", product)
	log.Println("Hasil storeRes ", storeRes)
	log.Println("hasil category res : ", categoryRes)
	// find -> categoryid
	// Save product with store id an category id

	return entities.ProductResponse{
		Id:          product.ID,
		Stock:       product.Stock,
		Barcode:     product.Barcode,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
		CategoryResponse: entities.CategoryResponse{
			Id:           categoryRes.ID,
			CategoryName: categoryRes.CategoryName,
			Description:  categoryRes.Description,
			ImageUrl:     categoryRes.ImageUrl,
			CreatedAt:    categoryRes.CreatedAt,
			UpdatedAt:    categoryRes.UpdatedAt,
		},
		StoreResponse: entities.StoreResponse{
			Id:        storeRes.ID,
			StoreName: storeRes.StoreName,
			Address:   storeRes.Address,
			OwnerName: storeRes.OwnerName,
			CreatedAt: storeRes.CreatedAt,
			UpdatedAt: storeRes.UpdatedAt,
		},
	}, nil
}

func (ps *productService) UpdateProduct(productReq entities.ProductRequest) (entities.ProductResponse, error) {

	// todo -> todo -> find by id product

	res, _ := ps.ProductRepository.FindById(productReq.Id)

	log.Println(res)

	return entities.ProductResponse{}, nil
}
