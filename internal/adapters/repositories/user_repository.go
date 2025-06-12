package repositories

import (
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(user *entities.User) error
		FindByEmail(email string) (*entities.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {

	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *entities.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) FindByEmail(email string) (*entities.User, error) {

	user := entities.User{}

	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}
