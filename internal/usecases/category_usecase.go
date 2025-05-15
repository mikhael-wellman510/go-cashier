package usecases

import (
	"log"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/entities"
)

type (
	CategoryService interface {
		CreateCategory(categoryReq entities.CategoryRequest) (entities.CategoryResponse, error)
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
