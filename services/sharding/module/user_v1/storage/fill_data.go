package storage

import (
	"context"
	"fmt"
	"shard/module/user_v1/domain"
)

func (s *postgresUserStorage) Fill(ctx context.Context,
) error {
	db := s.db
	for i := 0; i < 1000; i++ {
		db.Model(&domain.UserModel{}).Create(&domain.UserModel{
			UserID:   "User" + fmt.Sprint(i),
			Password: "Password" + fmt.Sprint(i),
		})
	}
	return nil
}
