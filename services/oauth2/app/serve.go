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
