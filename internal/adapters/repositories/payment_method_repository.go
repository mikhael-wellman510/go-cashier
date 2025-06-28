package repositories

import (
	"mikhael-project-go/internal/entities"

	"gorm.io/gorm"
)

type (
	PaymentMethodRepository interface {
		Create(paymentMethod *entities.PaymentMethod) error
		FindById(id string) (*entities.PaymentMethod, error)
		FindAll() ([]entities.PaymentMethod, error)
	}

	paymentMethodRepository struct {
		db *gorm.DB
	}
)

func NewPaymentMethodRepository(db *gorm.DB) PaymentMethodRepository {

	return &paymentMethodRepository{
		db: db,
	}
}

func (pmr *paymentMethodRepository) Create(paymentMethod *entities.PaymentMethod) error {

	return pmr.db.Create(paymentMethod).Error
}

func (pmr *paymentMethodRepository) FindById(id string) (*entities.PaymentMethod, error) {
	var paymentMethod entities.PaymentMethod

	err := pmr.db.First(&paymentMethod, "id=?", id).Error

	return &paymentMethod, err
}

func (pmr *paymentMethodRepository) FindAll() ([]entities.PaymentMethod, error) {
	var paymentMethod []entities.PaymentMethod

	err := pmr.db.Find(&paymentMethod).Error

	return paymentMethod, err
}
