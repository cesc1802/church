package handlers

import (
	"context"
	"log"
	"net"
	core "services.core-service"
	"services.rbac-service/module/auth_v1/transport/grpc_auth"
)

func ServeLoginServer(sc core.ServiceContext, addr string) (err error) {
	defer log.Printf("Login server stopped", err)
	s, err := grpc_auth.NewLoginServer(sc.context)

	ctx, cancel := context.WithCancel(ctx)
	errChan := make(chan error, 1)

	go func(ctx context.Context, errChan chan error) {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			errChan <- err
		}
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}(ctx, errChan)

	log.Printf("Login server started at %s\n", addr)

	select {
	case <-ctx.Done():
		cancel()
		return nil
	case err = <-errChan:
		cancel()
		return err
	}
}

func ServeRegisServer(ctx context.Context, addr string) (err error) {
	defer log.Printf("Register server stopped", err)

	s, err := NewRegisterServer()

	ctx, cancel := context.WithCancel(ctx)
	errChan := make(chan error, 1)

	go func(ctx context.Context, errChan chan error) {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			errChan <- err
		}
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}(ctx, errChan)

	log.Printf("Register server started at %s\n", addr)

	select {
	case <-ctx.Done():
		cancel()
		return nil
	case err = <-errChan:
		cancel()
		return err
	}
}
