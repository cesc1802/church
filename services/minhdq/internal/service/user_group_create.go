package service

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"minhdq/internal/model"
	"minhdq/internal/persistence"
)

type UserGroupCreateCommand struct {
	GroupID int `json:"group_id" valid:"required,int"`
	UserID  int `json:"user_id" valid:"required,int"`
}

func (c UserGroupCreateCommand) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func CreateUserGroup(ctx context.Context, c UserGroupCreateCommand) (err error) {
	err = c.Validate()

	if err != nil {
		return ErrInvalidInput
	}

	err = persistence.UserGroup().Save(ctx, model.UserGroup{
		GroupID:   c.GroupID,
		UserID:    c.UserID,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return ErrDatabase
	}

	return
}
