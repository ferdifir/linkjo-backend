package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type RegisterRequest struct {
	OwnerName  string  `json:"owner_name" validate:"required"`
	OutletName string  `json:"outlet_name" validate:"required"`
	Email      string  `json:"email" validate:"required,email"`
	Password   string  `json:"password" validate:"required,min=6"`
	Phone      string  `json:"phone" validate:"required"`
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	IsActive   bool    `json:"is_active"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateStatusRequest struct {
	IsActive bool `json:"is_active"`
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
