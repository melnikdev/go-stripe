package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Quantity int    `gorm:"not null"`
	StripeID string `gorm:"unique"`
	Prices   []Price
}
