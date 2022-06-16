package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"minhdq/internal/authentication"
	"minhdq/internal/persistence"
)

var loginServer *LoginServer

type LoginServer struct {
	authentication.UnimplementedLoginServer
}

func GetLoginServer() *LoginServer {
	if loginServer == nil {
		loginServer = &LoginServer{}
	}

	return loginServer
}

func (s *LoginServer) Login(ctx context.Context, in *authentication.LoginModel) (*authentication.JWT, error) {
	jwt, err := persistence.Account().Login(in.GetLoginID(), in.GetPassword())
	if err != nil {
		return nil, err
	}

	return &authentication.JWT{
		JWT:        jwt,
		Duratation: 24 * 2 * 60 * 60,
	}, nil
}

func (s *LoginServer) CheckAuthen(ctx context.Context, in *authentication.JWT) (*empty.Empty, error) {
	err := persistence.Account().Authentization(in.GetJWT())

	return &empty.Empty{}, err
}
