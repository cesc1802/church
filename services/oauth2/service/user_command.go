package service

import (
	"context"
	"oauth/model"
	"oauth/persistence"
)

func FindAllUser(ctx context.Context) (users []*model.User, err error) {
	return persistence.User().FindAll(ctx)
}

type UserCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cmd *UserCommand) Register(ctx context.Context) (err error) {
	return persistence.User().CreateUser(ctx, cmd.Username, cmd.Password)
}

func (cmd *UserCommand) FindByUserName(ctx context.Context) (user *model.User, err error) {
	return persistence.User().FindOneByID(ctx, cmd.Username)
}

func (cmd *UserCommand) VerifyUser(ctx context.Context) (err error) {
	return persistence.User().CheckUser(ctx, cmd.Username, cmd.Password)
}
