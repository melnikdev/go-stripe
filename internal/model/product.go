package model

type Product struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"unique"`
	Price    int    `gorm:"not null"`
	Quantity int    `gorm:"not null"`
}
