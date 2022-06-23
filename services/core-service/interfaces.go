package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gopkg.in/olahol/melody.v1"
)

type HasPrefix interface {
	Get() interface{}
	GetPrefix() string
}

type Storage interface {
	Get(prefix string) (interface{}, bool)
	MustGet(prefix string) interface{}
}

type Runnable interface {
	Name() string
	Configure() error
	Start() error
	Stop(ctx context.Context) error
}

type PrefixRunnable interface {
	HasPrefix
	Runnable
}

type HTTPServerHandler = func(*gin.Engine)

type GrpcServer interface {
	HasPrefix
	Runnable
	AddgRPCServer(s *grpc.Server)
}

type HttpServer interface {
	Runnable
	AddHandler(HTTPServerHandler)
	GetMelody() *melody.Melody
}

type ServiceContext interface {
	Get(prefix string) (interface{}, bool)
	MustGet(prefix string) interface{}
}

type Service interface {
	ServiceContext
	Name() string
	Version() string
	HttpServer() HttpServer
	GrpcServers() []GrpcServer
	GrpcServer(prefix string) GrpcServer
	Run() error
	Stop() error
}
