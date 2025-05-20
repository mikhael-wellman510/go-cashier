package usecases

import (
	"log"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/pkg/constants"
	"os"
	"strings"
)

type (
	CategoryService interface {
		CreateCategory(categoryReq entities.CategoryRequest) (entities.CategoryResponse, error)
		UpdateCategory(categoryreq entities.CategoryRequest) (entities.CategoryResponse, error)
	}

	categoryService struct {
		CategoryRepository repositories.CategoriesRepository
	}
)

func NewCategoryService(categoryRepository repositories.CategoriesRepository) CategoryService {
	return &categoryService{
		CategoryRepository: categoryRepository,
	}
}

func (cs *categoryService) CreateCategory(categoryReq entities.CategoryRequest) (entities.CategoryResponse, error) {

	category := entities.Categories{
		CategoryName: categoryReq.CategoryName,
		Description:  categoryReq.Description,
		ImageUrl:     categoryReq.ImageUrl,
	}

	if err := cs.CategoryRepository.Create(&category); err != nil {
		return entities.CategoryResponse{}, err
	}

	log.Println("User case : ", category)
	return entities.CategoryResponse{
		Id:           category.ID,
		CategoryName: category.CategoryName,
		Description:  category.Description,
		ImageUrl:     category.ImageUrl,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}, nil
}

func (cs *categoryService) UpdateCategory(categoryreq entities.CategoryRequest) (entities.CategoryResponse, error) {

	// todo find by id
	res, err := cs.CategoryRepository.FindById(categoryreq.Id)
	_ = os.Remove(strings.TrimPrefix(res.ImageUrl, constants.Server))
	if err != nil {

		return entities.CategoryResponse{}, err
	}

	// todo -> generate foto dulu  biar ke save

	res.CategoryName = categoryreq.CategoryName
	res.Description = categoryreq.Description
	res.ImageUrl = categoryreq.ImageUrl

	// updated
	if err := cs.CategoryRepository.Update(&res); err != nil {
		return entities.CategoryResponse{}, err
	}

	// Hapus data yg lama karena mau di updated

	return entities.CategoryResponse{
		Id:           res.ID,
		CategoryName: res.CategoryName,
		Description:  res.Description,
		ImageUrl:     res.ImageUrl,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}, nil
}
