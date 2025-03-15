package database

import (
	"github.com/melnikdev/go-stripe/internal/config"
	"github.com/melnikdev/go-stripe/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{}, &model.Product{}, &model.Price{})
}
