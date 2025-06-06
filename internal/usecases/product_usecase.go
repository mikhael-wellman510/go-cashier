package usecases

import (
	"log"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/utils"
)

type (
	ProductService interface {
		CreateProduct(productReq *entities.ProductRequest) (*entities.ProductResponse, error)
		UpdateProduct(productReq *entities.ProductRequest) (*entities.ProductResponse, error)
		FindProductById(id string) (*entities.ProductResponse, error)
		FilterAndPaggingProduct(page int, limit int, search string) (utils.PaginationResponse, error)
		ExportProductToCsv(date string) ([]entities.ProductResponse, error)
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

func (ps *productService) CreateProduct(productReq *entities.ProductRequest) (*entities.ProductResponse, error) {

	// find -> store id
	storeRes, err := ps.StoreService.FindById(productReq.StoreId)

	if err != nil {
		return nil, err
	}

	categoryRes, err := ps.CategoryService.FindById(productReq.CategoriesId)

	if err != nil {
		return nil, err
	}

	product := &entities.Product{
		Stock:        productReq.Stock,
		Barcode:      productReq.Barcode,
		ProductName:  productReq.ProductName,
		Description:  productReq.Description,
		Price:        productReq.Price,
		CategoriesId: categoryRes.ID,
		StoreId:      storeRes.ID,
	}

	if err := ps.ProductRepository.Create(product); err != nil {
		return nil, err
	}
	// find -> categoryid
	// Save product with store id an category id

	return &entities.ProductResponse{
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

func (ps *productService) UpdateProduct(productReq *entities.ProductRequest) (*entities.ProductResponse, error) {

	// todo -> todo -> find by id product

	res, err := ps.ProductRepository.FindById(productReq.Id)

	if err != nil {
		return nil, err
	}

	// Updated
	data := &entities.Product{
		Base: entities.Base{
			ID: res.ID,
			// CreatedAt: res.CreatedAt,
		},
		Stock:        productReq.Stock,
		Barcode:      productReq.Barcode,
		ProductName:  productReq.ProductName,
		Description:  productReq.Description,
		Price:        productReq.Price,
		CategoriesId: productReq.CategoriesId,
		StoreId:      productReq.StoreId,
	}

	if err := ps.ProductRepository.Update(data); err != nil {
		log.Println("Error ini ", err.Error())
		return nil, err
	}

	product, _ := ps.ProductRepository.FindById(productReq.Id)

	return &entities.ProductResponse{
		Id:          data.ID,
		Stock:       data.Stock,
		Barcode:     data.Barcode,
		ProductName: data.ProductName,
		Description: data.Description,
		Price:       data.Price,
		CategoryResponse: entities.CategoryResponse{
			Id:           product.Categories.ID,
			CategoryName: product.Categories.CategoryName,
			Description:  product.Categories.Description,
			ImageUrl:     product.Categories.ImageUrl,
			CreatedAt:    product.Categories.CreatedAt,
			UpdatedAt:    product.Categories.UpdatedAt,
		},
		StoreResponse: entities.StoreResponse{
			Id:        product.Store.ID,
			StoreName: product.Store.StoreName,
			Address:   product.Store.Address,
			OwnerName: product.Store.OwnerName,
			CreatedAt: product.Store.CreatedAt,
			UpdatedAt: product.Store.UpdatedAt,
		},
	}, nil

}
func (ps *productService) FindProductById(id string) (*entities.ProductResponse, error) {

	res, err := ps.ProductRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	log.Println("Hasil : ", res)

	return &entities.ProductResponse{
		Id:          res.ID,
		Stock:       res.Stock,
		Barcode:     res.Barcode,
		ProductName: res.ProductName,
		Description: res.Description,
		Price:       res.Price,
		CategoryResponse: entities.CategoryResponse{
			Id:           res.Categories.ID,
			CategoryName: res.Categories.CategoryName,
			Description:  res.Categories.Description,
			ImageUrl:     res.Categories.ImageUrl,
			CreatedAt:    res.Categories.CreatedAt,
			UpdatedAt:    res.Categories.UpdatedAt,
		},
		StoreResponse: entities.StoreResponse{
			Id:        res.Store.ID,
			StoreName: res.Store.StoreName,
			Address:   res.Store.Address,
			OwnerName: res.Store.OwnerName,
			CreatedAt: res.Store.CreatedAt,
			UpdatedAt: res.Store.UpdatedAt,
		},
	}, nil
}

func (ps *productService) FilterAndPaggingProduct(page int, limit int, search string) (utils.PaginationResponse, error) {

	res, err := ps.ProductRepository.FindAllPagging(page, limit, search)

	if err != nil {
		return utils.PaginationResponse{}, err
	}
	// Todo -> total limit ketika di filter tanpa limit dan offset
	count, err := ps.ProductRepository.CountWithFilterProduct(search)

	if err != nil {
		return utils.PaginationResponse{}, err
	}

	result := utils.PaginationDto(res, int(count), page, limit)
	return result, nil
}

func (ps *productService) ExportProductToCsv(date string) ([]entities.ProductResponse, error) {

	res, err := ps.ProductRepository.FindProductByDate(date)
	if err != nil {
		return []entities.ProductResponse{}, err
	}

	response := []entities.ProductResponse{}
	for _, product := range res {
		// Find store
		store, err := ps.StoreService.FindById(product.StoreId)
		if err != nil {
			return []entities.ProductResponse{}, err
		}
		category, err := ps.CategoryService.FindById(product.CategoriesId)

		if err != nil {
			return []entities.ProductResponse{}, err
		}
		log.Println("Hasil dari store : ", store)
		response = append(response, entities.ProductResponse{
			Id:          product.ID,
			Stock:       product.Stock,
			Barcode:     product.Barcode,
			ProductName: product.ProductName,
			Description: product.Description,
			Price:       product.Price,
			CategoryResponse: entities.CategoryResponse{
				Id:           category.ID,
				CategoryName: category.CategoryName,
				Description:  category.Description,
				ImageUrl:     category.ImageUrl,
				CreatedAt:    category.CreatedAt,
				UpdatedAt:    category.UpdatedAt,
			},
			StoreResponse: entities.StoreResponse{
				Id:        store.ID,
				StoreName: store.StoreName,
				Address:   store.Address,
				OwnerName: store.OwnerName,
				CreatedAt: store.CreatedAt,
				UpdatedAt: store.UpdatedAt,
			},
		})

	}

	return response, nil
}
