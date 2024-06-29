package handlers

import (
	"context"
	_ "fiap-tech-challenge-producao/docs"
	"fiap-tech-challenge-producao/internal/adapters/handlers/http"
	"fiap-tech-challenge-producao/internal/adapters/handlers/pubsub"
	"fmt"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/messaging"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"os"
)

type Server struct {
	appName               *string
	host                  string
	Server                *echo.Echo
	healthHandler         *http.HealthCheck
	producaoHandler       *http.Producao
	producaoPubsubHandler *pubsub.ProducaoHandler
	messageClient         messaging.Client
}

// NewAPIServer creates the main http with all configurations necessary
func NewAPIServer(healthHandler *http.HealthCheck, producaoHandler *http.Producao, producaoPubsubHandler *pubsub.ProducaoHandler, messageClient messaging.Client) *Server {
	host := os.Getenv("SERVER_PORT")
	if host == "" {
		host = ":3000"
	}

	appName := "tech-challenge-producao"
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
		appName:               &appName,
		host:                  host,
		Server:                app,
		healthHandler:         healthHandler,
		producaoHandler:       producaoHandler,
		producaoPubsubHandler: producaoPubsubHandler,
		messageClient:         messageClient,
	}
}

func (hs *Server) RegisterHandlers() {
	hs.healthHandler.RegisterHealth(hs.Server)
	hs.producaoHandler.RegistraRotasFila(hs.Server)
}

// Start starts an application on specific port
func (hs *Server) Start(ctx context.Context) {
	producaoQueue := os.Getenv("PRODUCAO_QUEUE")
	if producaoQueue == "" {
		log.Fatalf("env variable PRODUCAO_QUEUE not set")
	}
	//adiciona listener
	go hs.messageClient.Listen(ctx, producaoQueue, hs.producaoPubsubHandler.AdicionaProducaoHandler)

	hs.RegisterHandlers()
	log.Info(ctx, fmt.Sprintf("Starting a http at http://%s", hs.host))
	err := hs.Server.Start(hs.host)
	if err != nil {
		log.Error(ctx, errorx.Decorate(err, "failed to start the http server"))
		return
	}
}
