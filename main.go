package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
)

//go:generate wire

//	@title			Tech Challenge API
//	@version		3.0.0
//	@description	This is a documentation of all endpoints in the API.

// @host		localhost:3000
// @BasePath	/
// @schemes http
// @produce json
// @securityDefinitions.apikey	JWT
// @in							header
// @name						token
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app, err := InitializeWebServer()
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	app.Start(ctx)

	// listens for system signals to gracefully shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	switch <-signalChannel {
	case os.Interrupt:
		log.Info("Received Interrupt, stopping...")
		cancel()
	case syscall.SIGTERM:
		log.Info("Received SIGTERM, stopping...")
		cancel()
	case syscall.SIGINT:
		log.Info("Received SIGINT, stopping...")
		cancel()
	}
}
