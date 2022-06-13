package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"minhdq/internal/authentication"
	"minhdq/internal/persistence"
)

var registServer *RegisterServer

type RegisterServer struct {
	authentication.UnimplementedResgisterServer
}

func GetRegisServer() *RegisterServer {
	if registServer == nil {
		registServer = &RegisterServer{}
	}

	return registServer
}

func (s *RegisterServer) Resgister(ctx context.Context, in *authentication.RegisterModel) (*empty.Empty, error) {
	if in.GetLoginID() == "" {
		return nil, errors.New("LoginID Can't be null")
	}

	accounts := persistence.Account()

	fmt.Println(in.String())
	err := accounts.Register(in.GetLoginID(), in.GetPassword(), in.GetFirstName(), in.LastName)

	return &empty.Empty{}, err
}
