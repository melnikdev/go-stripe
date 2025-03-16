package service

import "gorm.io/gorm"

type SubscriptionService interface {
	CreateSubscription(customerID, priceID string)
}

type subscriptionService struct {
	db *gorm.DB
}

func NewSubscriptionService(db *gorm.DB) SubscriptionService {
	return &subscriptionService{
		db: db,
	}
}

func (s *subscriptionService) CreateSubscription(customerID, priceID string) {

}
