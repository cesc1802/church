package business

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"services.core-service/app_error"
	"services.rbac-service/errorcode"
	"services.rbac-service/module/auth_v1/dto"
	"services.rbac-service/module/user_v1/domain"
	"time"
)

type LoginStorage interface {
	FindOneByLoginID(ctx context.Context, loginID string) (*domain.UserModel, error)
	FindByConditions(ctx context.Context,
		conditions map[string]interface{}) (*domain.UserModel, error)
}

type LoginBusiness struct {
	store LoginStorage
}

func NewLoginBusiness(store LoginStorage) *LoginBusiness {
	return &LoginBusiness{
		store: store,
	}
}

func (biz *LoginBusiness) UserLogin(ctx context.Context,
	input *dto.LoginRequest) (*dto.LoginResponse, error) {

	conditions := map[string]interface{}{"LoginID": input.LoginID, "Password": input.Password}

	user, err := biz.store.FindByConditions(ctx, conditions)

	if err != nil {
		return nil, app_error.NewCustomError(err, "", errorcode.ErrCannotGetUser)
	}

	if user.Status == 0 {
		return nil, app_error.NewCustomError(nil, "", errorcode.ErrUserHasBeenBlock)
	}

	if err != nil {
		return nil, app_error.NewCustomError(err, "", errorcode.ErrCannotCreateUser)
	}

	expr := 24 * 2 * 60 * 60

	j := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        user.LoginID,
		"IssuedAt":  time.Now().Unix(),
		"Issuer":    "minhdq",
		"ExpiresAt": expr,
	})

	token, err := j.SignedString([]byte("SECRET"))

	if err != nil {
		return nil, app_error.NewCustomError(err, "", errorcode.ErrCannotCreateUser)
	}

	return &dto.LoginResponse{Token: token}, nil
}
