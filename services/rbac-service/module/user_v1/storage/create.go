package storage

import (
	"context"

	"gorm.io/gorm"
	"services.rbac-service/module/user_v1/domain"
)

func (s *postgresUserStorage) Create(ctx context.Context, input *domain.CreateUserModel) error {
	db := s.db

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(input.TableName()).Create(input).Error; err != nil {
			return err
		}
		return nil
	})
}
