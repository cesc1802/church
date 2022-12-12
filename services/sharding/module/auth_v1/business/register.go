package business

import (
	"context"
	"shard/errorcode"
	"shard/module/auth_v1/dto"
	"shard/module/user_v1/domain"

	"services.core-service/app_error"
)

type RegisterStorage interface {
	FindOneByLoginID(ctx context.Context, loginID string) (*domain.UserModel, error)
	Create(ctx context.Context, input *domain.CreateUserModel) error
	ListAll(ctx context.Context) (*[]domain.UserModel, error)
}

type RegisterBusiness struct {
	store RegisterStorage
}

func (biz *RegisterBusiness) FindOneByLoginID(ctx context.Context, loginID string) (*domain.UserModel, error) {
	user, err := biz.store.FindOneByLoginID(ctx, loginID)
	if err != nil {
		return nil, app_error.NewCustomError(err, "", errorcode.ErrCannotGetUser)
	}
	return user, nil
}

func (biz *RegisterBusiness) Create(ctx context.Context, input *dto.RegisterRequest) error {
	user, err := biz.store.FindOneByLoginID(ctx, input.UserID)
	if err != nil {
		return app_error.NewCustomError(err, "", errorcode.ErrCannotGetUser)
	}

	if user.Status == 0 {
		return app_error.NewCustomError(nil, "", errorcode.ErrUserHasBeenBlock)
	}

	err = biz.store.Create(ctx, &domain.CreateUserModel{
		UserID:   input.UserID,
		Password: input.Password,
	},
	)

	if err != nil {
		return app_error.NewCustomError(err, "", errorcode.ErrCannotCreateUser)
	}

	if err != nil {
		return app_error.NewCustomError(err, "", errorcode.ErrCannotGetUser)
	}

	if user.Status == 0 {
		return app_error.NewCustomError(nil, "", errorcode.ErrUserHasBeenBlock)
	}

	err = biz.store.Create(ctx, &domain.CreateUserModel{
		UserID:   input.UserID,
		Password: input.Password,
	},
	)

	if err != nil {
		return app_error.NewCustomError(err, "", errorcode.ErrCannotCreateUser)
	}

	return nil
}

func (biz *RegisterBusiness) ListAll(ctx context.Context) (*[]domain.UserModel, error) {
	return biz.store.ListAll(ctx)
}

func NewRegisterBusiness(store RegisterStorage) *RegisterBusiness {
	return &RegisterBusiness{
		store: store,
	}
}
