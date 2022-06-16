package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
)

func Serve(ctx context.Context, addr string) (err error) {
	defer log.Printf("HTTP server stopped", err)

	r := NewChiHandeler()

	ctx, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}

	srv := http.Server{
		Addr:    addr,
		Handler: r,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	errChan := make(chan error, 1)

	go func(ctx context.Context, errChan chan error) {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}(ctx, errChan)

	log.Printf("HTTP server started at %s\n", addr)

	select {
	case <-ctx.Done():
		cancel()
		wg.Wait()
		return nil
	case err = <-errChan:
		cancel()
		return err
	}
}

func ServeLoginServer(ctx context.Context, addr string) (err error) {
	defer log.Printf("Login server stopped", err)

	s, err := NewLoginServer()

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
