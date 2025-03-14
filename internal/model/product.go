package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Price    int    `gorm:"not null"`
	Quantity int    `gorm:"not null"`
}
