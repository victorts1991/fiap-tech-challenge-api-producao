package http

import (
	"context"
	_ "fiap-tech-challenge-api/docs"
	"fiap-tech-challenge-api/internal/adapters/http/handlers"
	"fmt"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

type Server struct {
	appName         *string
	host            string
	Server          *echo.Echo
	healthHandler   *handlers.HealthCheck
	producaoHandler *handlers.Producao
}

// NewAPIServer creates the main http with all configurations necessary
func NewAPIServer(healthHandler *handlers.HealthCheck, pedidoHandler *handlers.Producao) *Server {
	host := ":3003"
	appName := "tech-challenge-api"
	app := echo.New()

	app.HideBanner = true
	app.HidePort = true

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info(
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	app.GET("/docs/*", echoSwagger.WrapHandler)

	return &Server{
		appName:       &appName,
		host:          host,
		Server:        app,
		healthHandler: healthHandler,
		//clienteHandler:   clienteHandler,
		//produtoHandler:   produtoHandler,
		producaoHandler: pedidoHandler,
		//pagamentoHandler: pagamentoHandler,
		//loginHandler:     loginHandler,
	}
}

func (hs *Server) RegisterHandlers() {
	hs.healthHandler.RegisterHealth(hs.Server)
	hs.producaoHandler.RegistraRotasFila(hs.Server)
}

// Start starts an application on specific port
func (hs *Server) Start() {
	hs.RegisterHandlers()
	ctx := context.Background()
	log.Info(ctx, fmt.Sprintf("Starting a http at http://%s", hs.host))
	err := hs.Server.Start(hs.host)
	if err != nil {
		log.Error(ctx, errorx.Decorate(err, "failed to start the http server"))
		return
	}
}
