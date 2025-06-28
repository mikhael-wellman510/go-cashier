package entities

import "time"

type PaymentMethod struct {
	Base
	PaymentMethod string `json:"payment_method" gorm:"column:payment_method;not null;unique"`
}

type (
	PaymentMethodRequest struct {
		Id            string `json:"id"`
		PaymentMethod string `json:"payment_method"`
	}

	PaymentMethodResponse struct {
		Id            string    `json:"id"`
		PaymentMethod string    `json:"payment_method"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
	}
)
