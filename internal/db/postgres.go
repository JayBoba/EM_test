package db

import (
	"github.com/JayBoba/EM_test/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Subscription{}); err != nil {
		return nil, err
	}
	return db, nil
}
