package usecases

import (
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/entities"
)

type (
	PaymentMethodService interface {
		CreatePaymentMethod(paymentMethod *entities.PaymentMethodRequest) (*entities.PaymentMethodResponse, error)
		FindAllPaymentMethod() ([]entities.PaymentMethodResponse, error)
	}

	paymentMethodService struct {
		PaymentMethodRepository repositories.PaymentMethodRepository
	}
)

func NewPaymenMethodService(paymentMethodRepository repositories.PaymentMethodRepository) PaymentMethodService {

	return &paymentMethodService{
		PaymentMethodRepository: paymentMethodRepository,
	}
}

func (pmr *paymentMethodService) CreatePaymentMethod(paymentMethod *entities.PaymentMethodRequest) (*entities.PaymentMethodResponse, error) {

	paymentMethodData := &entities.PaymentMethod{
		PaymentMethod: paymentMethod.PaymentMethod,
	}

	if err := pmr.PaymentMethodRepository.Create(paymentMethodData); err != nil {
		return nil, err
	}

	return &entities.PaymentMethodResponse{
		Id:            paymentMethod.Id,
		PaymentMethod: paymentMethod.PaymentMethod,
		CreatedAt:     paymentMethodData.CreatedAt,
		UpdatedAt:     paymentMethodData.UpdatedAt,
	}, nil
}

func (pmr *paymentMethodService) FindAllPaymentMethod() ([]entities.PaymentMethodResponse, error) {

	res, err := pmr.PaymentMethodRepository.FindAll()

	if err != nil {
		return nil, err
	}

	paymentMethodResponse := []entities.PaymentMethodResponse{}

	for _, data := range res {
		payResponse := entities.PaymentMethodResponse{
			Id:            data.ID,
			PaymentMethod: data.PaymentMethod,
			CreatedAt:     data.CreatedAt,
			UpdatedAt:     data.UpdatedAt,
		}
		paymentMethodResponse = append(paymentMethodResponse, payResponse)
	}

	return paymentMethodResponse, nil
}
