package grpc_auth

import (
	"context"

	"services.rbac-service/module/auth_v1/business"
	"services.rbac-service/module/auth_v1/dto"
	"services.rbac-service/module/user_v1/grpc_user"
)

type ILoginBusiness interface {
	UserLogin(ctx context.Context,
		input *dto.LoginRequest,
	) (*dto.LoginResponse, error)
}

type LoginBusinessgRPC struct {
	grpc_user.UnimplementedLoginServer
	ILoginBusiness
}

func NewLogingRPCBusiness(store business.LoginStorage) *LoginBusinessgRPC {
	lbusiness := business.NewLoginBusiness(store)
	return &LoginBusinessgRPC{
		ILoginBusiness: lbusiness,
	}
}

func (biz *LoginBusinessgRPC) Login(ctx context.Context, in *grpc_user.LoginModel) (*grpc_user.JWT, error) {
	jwt, err := biz.UserLogin(ctx, &dto.LoginRequest{
		LoginID:  in.GetLoginID(),
		Password: in.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &grpc_user.JWT{
		Token: jwt.Token,
	}, nil
}
