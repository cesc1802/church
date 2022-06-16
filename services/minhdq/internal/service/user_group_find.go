package service

import (
	"context"
	"fmt"

	"minhdq/internal/model"
	"minhdq/internal/persistence"
)

func UserGroupGetAll(ctx context.Context) (data []*model.UserGroup, err error) {
	data, err = persistence.UserGroup().FindAll(ctx)
	fmt.Println(err)
	if err != nil {
		return nil, ErrDatabase
	}

	return
}
