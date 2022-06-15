package service

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"minhdq/internal/model"
	"minhdq/internal/persistence"
)

type UserGroupUpdateCommand struct {
	GroupID   int       `json:"group_id" valid:"required,int"`
	UserID    int       `json:"user_id" valid:"required,int"`
	CreatedAt time.Time `json:"created_at" valid:"required,"`
}

func (c UserGroupUpdateCommand) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func UpdateUserGroup(ctx context.Context, c UserGroupUpdateCommand) (data model.UserGroup, err error) {
	err = c.Validate()

	if err != nil {
		return data, ErrInvalidInput
	}

	data, err = persistence.UserGroup().Update(ctx, model.UserGroup{
		GroupID:   c.GroupID,
		UserID:    c.UserID,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return data, ErrDatabase
	}

	return
}
