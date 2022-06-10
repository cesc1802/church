package service

import (
	"context"
	"github.com/asaskevich/govalidator"
	"minhdq/internal/persistence"
)

type UserGroupDeleteCommand struct {
	GroupID int `json:"group_id" valid:"required,int"`
	UserID  int `json:"user_id" valid:"required,int"`
}

func (c UserGroupDeleteCommand) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func DeleteUserGroup(ctx context.Context, c UserGroupDeleteCommand) (err error) {
	err = c.Validate()

	if err != nil {
		return ErrInvalidInput
	}

	err = persistence.UserGroup().Delete(ctx, c.UserID, c.GroupID)

	if err != nil {
		return ErrDatabase
	}

	return
}
