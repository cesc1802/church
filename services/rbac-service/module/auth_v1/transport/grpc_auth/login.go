package grpc_auth

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
	core "services.core-service"
	"services.rbac-service/constants"
	"services.rbac-service/module/user_v1/grpc_user"
	"services.rbac-service/module/user_v1/storage"
)

func NewLoginServer(sc core.ServiceContext) (*grpc.Server, error) {

	db := sc.MustGet(constants.KeyMainDb).(*gorm.DB)
	store := storage.NewPostgresUserStorage(db)

	s := grpc.NewServer()
	loginS := NewLogingRPCBusiness(store)
	grpc_user.RegisterLoginServer(s, loginS)

	return s, nil
}
