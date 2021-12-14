package business

import (
	"context"
	"services.core-service/app_error"
	"services.rbac-service/errorcode"
	"services.rbac-service/module/auth_v1/dto"
	"services.rbac-service/module/user_v1/domain"
)

type RegisterStorage interface {
	FindOneByLoginID(ctx context.Context, loginID string) (*domain.UserModel, error)
	Create(ctx context.Context, input *domain.CreateUserModel) error
}

type registerBusiness struct {
	store RegisterStorage
}

func NewRegisterBusiness(store RegisterStorage) *registerBusiness {
	return &registerBusiness{
		store: store,
	}
}

func (biz *registerBusiness) UserRegister(ctx context.Context,
	input *dto.RegisterRequest) error {
	user, err := biz.store.FindOneByLoginID(ctx, input.LoginID)

	if err != nil {
		return app_error.NewCustomError(err, "", errorcode.ErrCannotGetUser)
	}

	if user.Status == 0 {
		return app_error.NewCustomError(nil, "", errorcode.ErrUserHasBeenBlock)
	}

	err = biz.store.Create(ctx, &domain.CreateUserModel{LoginID: input.LoginID,
		Password:  input.Password,
		LastName:  input.LastName,
		FirstName: input.FirstName},
	)

	if err != nil {
		return app_error.NewCustomError(err, "", errorcode.ErrCannotCreateUser)
	}

	return nil
}
