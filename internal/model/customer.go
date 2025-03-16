package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique"`
	StripeID string `gorm:"unique"`
}
