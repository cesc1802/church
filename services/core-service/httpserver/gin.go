package httpserver

import (
	"context"
	"fmt"
	"gopkg.in/olahol/melody.v1"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	config "services.core-service/configs"
	"services.core-service/httpserver/middleware"
	"services.core-service/i18n"
	"services.core-service/logger"
	baseValidation "services.core-service/validation"
)

type ginOpt struct {
	name string
	port string
	host string
	mode string
}

type ginService struct {
	isRunning bool
	engine    *gin.Engine
	*http.Server
	handlers []func(engine *gin.Engine)
	i18n     *i18n.I18n
	melody   *melody.Melody
	*ginOpt
}

func (gs *ginService) GetMelody() *melody.Melody {
	return gs.melody
}

func New(c config.Config, i18n *i18n.I18n) *ginService {
	return &ginService{
		isRunning: false,
		i18n:      i18n,
		handlers:  []func(*gin.Engine){},
		melody:    melody.New(),
		ginOpt: &ginOpt{
			name: "GIN-SERVICE",
			port: c.ServerConfig.Port,
			host: c.ServerConfig.Host,
		},
	}
}

func (gs *ginService) Configure() error {
	if gs.isRunning {
		return nil
	}
	if gs.mode == "" {
		gs.mode = "debug"
	}

	gin.SetMode(gs.mode)
	gs.engine = gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(baseValidation.JsonTagNameFunc)
	}

	gs.engine.RedirectTrailingSlash = true
	gs.engine.RedirectFixedPath = true

	// Recovery
	// TODO: you can add more middleware here
	gs.engine.Use(middleware.Recovery(gs.i18n))

	gs.isRunning = true
	return nil
}

func (gs *ginService) Name() string {
	return gs.name
}

func (gs *ginService) Start() error {
	if err := gs.Configure(); err != nil {
		return err
	}

	// Setup handlers
	for _, hdl := range gs.handlers {
		hdl(gs.engine)
	}

	gs.Server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", gs.host, gs.port),
		Handler: gs.engine,
	}
	logger.Info("Listening and serving HTTP on %v:%v", gs.host, gs.port)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", gs.host, gs.port))
	if err != nil {
		logger.Info("Listening error: %v", err)
		return err
	}

	err = gs.Serve(lis)

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (gs *ginService) Stop(ctx context.Context) error {
	// ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancelFn()

	if gs.Server != nil {
		logger.Info("server shutting down....")
		_ = gs.Shutdown(ctx)
	}
	return nil
}

func (gs *ginService) AddHandler(hdl func(engine *gin.Engine)) {
	gs.handlers = append(gs.handlers, hdl)
}
