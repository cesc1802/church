package service

import (
	"context"
	"minhdq/internal/model"
	"minhdq/internal/persistence"
)

func UserGroupGetAll(ctx context.Context) (data []*model.UserGroup, err error) {
	data, err = persistence.UserGroup().FindAll(ctx)
	if err != nil {
		return nil, ErrDatabase
	}

	return
}
