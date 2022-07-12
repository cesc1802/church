package storage

import (
	"context"
	"shard/module/user_v1/domain"

	"gorm.io/gorm"
)

func (s *postgresUserStorage) Create(ctx context.Context, input *domain.CreateUserModel) error {
	db := s.db

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(domain.CreateUserModel{}).Create(input).Error; err != nil {
			return err
		}
		return nil
	})
}
