package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/melnikdev/go-stripe/internal/model"
	"github.com/melnikdev/go-stripe/internal/request"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(request request.CreateProductRequest) (*model.Product, error)
	UpdateStripeId(product *model.Product, stripeID string) (*model.Product, error)
	GetAll() ([]model.Product, error)
	GetById(id string) (*model.Product, error)
	CreatePrice(productId uint, priceId string, amount int) (*model.Price, error)
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

func (s *productService) Create(request request.CreateProductRequest) (*model.Product, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return nil, err
	}

	product := model.Product{
		Name:     request.Name,
		Quantity: request.Quantity,
	}

	result := s.db.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (s *productService) UpdateStripeId(product *model.Product, stripeID string) (*model.Product, error) {

	result := s.db.Model(&product).Update("stripe_id", stripeID)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (s *productService) GetAll() ([]model.Product, error) {
	var products []model.Product
	result := s.db.Where("stripe_id IS NOT NULL").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (s *productService) GetById(id string) (*model.Product, error) {
	var product model.Product
	result := s.db.Where("id = ?", id).Where("stripe_id IS NOT NULL").Preload("Prices").First(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (s *productService) CreatePrice(productId uint, priceId string, amount int) (*model.Price, error) {
	price := model.Price{
		ProductID: productId,
		PriceID:   priceId,
		Amount:    amount,
		Currency:  "usd",
		Type:      "recurring",
	}

	result := s.db.Create(&price)
	if result.Error != nil {
		return nil, result.Error
	}

	return &price, nil
}
