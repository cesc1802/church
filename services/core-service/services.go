package core

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type AppServiceOption func(service *AppService)

type AppService struct {
	name         string
	version      string
	ctx          context.Context
	cancel       func()
	signals      []os.Signal
	httpserver   HttpServer
	subServices  []Runnable
	initServices map[string]PrefixRunnable
}

func WithVersion(version string) AppServiceOption {
	return func(app *AppService) {
		app.version = version
	}
}

func WithName(name string) AppServiceOption {
	return func(app *AppService) {
		app.name = name
	}
}

func WithHttpServer(server HttpServer) AppServiceOption {
	return func(app *AppService) {
		app.httpserver = server
	}
}

func WithInitRunnable(runnable PrefixRunnable) AppServiceOption {
	return func(app *AppService) {
		app.subServices = append(app.subServices, runnable)
		app.initServices[runnable.GetPrefix()] = runnable
	}
}

func NewAppService(opts ...AppServiceOption) *AppService {
	sv := &AppService{
		signals:      []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		subServices:  []Runnable{},
		initServices: map[string]PrefixRunnable{},
	}

	if sv.ctx == nil {
		sv.ctx = context.Background()
	}

	// Setup cancelCtx to listen Ctrl + C
	cancelCtx, cancelFunc := context.WithCancel(sv.ctx)
	sv.ctx = cancelCtx
	sv.cancel = cancelFunc

	for _, opt := range opts {
		opt(sv)
	}
	if sv.httpserver != nil {
		sv.subServices = append(sv.subServices, sv.httpserver)
	}

	return sv
}

func (s *AppService) Run() error {
	eg, ctx := errgroup.WithContext(s.ctx)

	wg := sync.WaitGroup{}
	for _, sub := range s.subServices {
		srv := sub
		eg.Go(func() error {
			<-ctx.Done()
			sctx, cancel := context.WithTimeout(NewContext(context.Background(), srv), 5*time.Second)
			defer cancel()
			return srv.Stop(sctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start()
		})
	}
	wg.Wait()
	c := make(chan os.Signal, 1)
	signal.Notify(c, s.signals...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				_ = s.Stop()
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (s *AppService) Stop() error {
	if s.cancel != nil {
		s.cancel()
	}
	return nil
}

func (s *AppService) HttpServer() HttpServer {
	return s.httpserver
}

func (s *AppService) Get(prefix string) (interface{}, bool) {
	if prefixRunnable, ok := s.initServices[prefix]; ok {
		return prefixRunnable.Get(), true
	}
	return nil, false
}

func (s *AppService) MustGet(prefix string) interface{} {
	return s.initServices[prefix].Get()
}

func (s *AppService) Name() string {
	return s.name
}

func (s *AppService) Version() string {
	return s.version
}

type appServiceKey struct {
}

func NewContext(ctx context.Context, ra Runnable) context.Context {
	return context.WithValue(ctx, appServiceKey{}, ra)
}

func FromContext(ctx context.Context) (Runnable, bool) {
	ra, ok := ctx.Value(appServiceKey{}).(Runnable)
	return ra, ok
}
