package grpc_auth

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"services.rbac-service/module/auth_v1/business"
	"services.rbac-service/module/auth_v1/dto"
	"services.rbac-service/module/user_v1/grpc_user"
)

type RegisterBusinessgRPC struct {
	grpc_user.UnimplementedResgisterServer
	business.RegisterBusiness
}

func NewRegistergRPCBusiness(store business.RegisterStorage) *RegisterBusinessgRPC {
	return &RegisterBusinessgRPC{
		RegisterBusiness: *business.NewRegisterBusiness(store),
	}
}

func (biz *RegisterBusinessgRPC) Resgister(ctx context.Context, in *grpc_user.RegisterModel) (*empty.Empty, error) {
	if in.GetLoginID() == "" {
		return nil, errors.New("LoginID Can't be null")
	}

	err := biz.UserRegister(ctx, &dto.RegisterRequest{
		LoginID:   in.GetLoginID(),
		Password:  in.GetPassword(),
		LastName:  &in.LastName,
		FirstName: &in.FirstName,
	})

	return &empty.Empty{}, err
}
