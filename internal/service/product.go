package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/melnikdev/go-stripe/internal/model"
	"github.com/melnikdev/go-stripe/internal/request"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(request request.CreateProductRequest) (model.Product, error)
}

type productService struct {
	db       *gorm.DB
	validate *validator.Validate
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{
		db:       db,
		validate: validator.New(),
	}
}

func (s *productService) Create(request request.CreateProductRequest) (model.Product, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return model.Product{}, err
	}

	product := model.Product{
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	result := s.db.Create(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}

	return product, nil
}
