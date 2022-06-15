package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"

	config "services.core-service/configs"
	"services.core-service/i18n"
	"services.core-service/logger"
)

type gRPCOption struct {
	name string
	port string
	host string
	mode string
}

type gRPCService struct {
	isRunning bool
	prefix    string
	Server    *grpc.Server
	i18n      *i18n.I18n
	*gRPCOption
}

func (gs *gRPCService) Get() interface{} {
	return gs.name
}

func (gs *gRPCService) GetPrefix() string {
	return gs.prefix
}

func (gs *gRPCService) Configure() error {
	// TODO implement me
	panic("implement me")
}

func (gs *gRPCService) Start() error {
	if gs.Server == nil {
		return errors.New("Server not init yet")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", gs.host, gs.port))
	if err != nil {
		logger.Info("Listening error: %v", err)
		return err
	}

	err = gs.Server.Serve(lis)

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (gs *gRPCService) Stop(ctx context.Context) error {
	if gs.Server != nil {
		logger.Info("Server shutting down....")
		_ = gs.Stop(ctx)
	}
	return nil
}

func New(c config.Config, i18n *i18n.I18n, prefix string) *gRPCService {
	return &gRPCService{
		isRunning: false,
		i18n:      i18n,
		prefix:    prefix,
		gRPCOption: &gRPCOption{
			name: "grpc-SERVICE",
			port: c.ServerConfig.Port,
			host: c.ServerConfig.Host,
		},
	}
}

func (gs *gRPCService) Name() string {
	return gs.name
}

func (gs *gRPCService) AddgRPCServer(s *grpc.Server) {
	gs.Server = s
}
