package grpc_auth

import (
	"context"
	"services.rbac-service/module/auth_v1/business"
	"services.rbac-service/module/auth_v1/dto"
	"services.rbac-service/module/user_v1/grpc_user"
)

type LoginBusinessgRPC struct {
	grpc_user.UnimplementedLoginServer
	business.LoginBusiness
}

func NewLogingRPCBusiness(store business.LoginStorage) *LoginBusinessgRPC {
	return &LoginBusinessgRPC{
		LoginBusiness: *business.NewLoginBusiness(store),
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
