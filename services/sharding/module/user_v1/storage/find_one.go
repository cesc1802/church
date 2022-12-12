package storage

import (
	"context"
	"shard/module/user_v1/domain"

	"gorm.io/gorm"
	"services.core-service/app_error"
)

func (s *postgresUserStorage) FindByConditions(ctx context.Context,
	conditions map[string]interface{},
) (*domain.UserModel, error) {
	db := s.db
	var user domain.UserModel
	if err := db.Table(domain.UserModel{}.TableName()).Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, app_error.RecordNotFound
		}
		return nil, app_error.ErrDBQuery(err)
	}
	return &user, nil
}

func (s *postgresUserStorage) FindOneByID(ctx context.Context, id uint64) (*domain.UserModel, error) {
	db := s.db
	var user domain.UserModel

	conditions := map[string]interface{}{
		"id": id,
	}

	db = db.Where(conditions)

	if err := db.Table(user.TableName()).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, app_error.RecordNotFound
		}
		return nil, app_error.ErrDBQuery(err)
	}

	return &user, nil
}

func (s *postgresUserStorage) FindOneByLoginID(ctx context.Context, loginID string) (*domain.UserModel, error) {
	db := s.db
	var user domain.UserModel

	conditions := map[string]interface{}{
		"login_id": loginID,
	}

	db = db.Where(conditions)

	if err := db.Table(user.TableName()).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, app_error.RecordNotFound
		}
		return nil, app_error.ErrDBQuery(err)
	}

	return &user, nil
}
