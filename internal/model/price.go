package model

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	ProductID uint   `gorm:"not null"`
	PriceID   string `gorm:"unique"`
	Amount    int    `gorm:"not null"`
	Currency  string
	Type      string
}
