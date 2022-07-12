package storage

import (
	"context"
	"gorm.io/gorm"
	"services.core-service/app_error"
	"shard/module/user_v1/domain"
)

func (s *postgresUserStorage) ListAll(ctx context.Context,
) (*[]domain.UserModel, error) {
	db := s.db
	var users []domain.UserModel
	if err := db.Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, app_error.RecordNotFound
		}
		return nil, app_error.ErrDBQuery(err)
	}
	return &users, nil
}
