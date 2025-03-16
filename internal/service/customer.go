package service

import (
	"github.com/melnikdev/go-stripe/internal/model"
	"github.com/melnikdev/go-stripe/internal/request"
	"gorm.io/gorm"
)

type CustomerService interface {
	CreateCustomer(request request.CreateCustomerRequest, StripeID string) (*model.Customer, error)
}

type customerService struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) CustomerService {
	return &customerService{
		db: db,
	}
}

func (s *customerService) CreateCustomer(request request.CreateCustomerRequest, StripeID string) (*model.Customer, error) {

	customer := model.Customer{
		Email:    request.Email,
		Name:     request.Name,
		StripeID: StripeID,
	}

	result := s.db.Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}
