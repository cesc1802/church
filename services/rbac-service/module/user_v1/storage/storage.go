package storage

import "gorm.io/gorm"

type postgresUserStorage struct {
	db *gorm.DB
}

func NewPostgresUserStorage(db *gorm.DB) *postgresUserStorage {
	return &postgresUserStorage{db: db}
}
