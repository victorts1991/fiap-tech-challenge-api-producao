package main

import (
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
	app, err := InitializeWebServer()
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	app.Start()

	// listens for system signals to gracefully shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		log.Info("Received SIGINT, stopping...")
	case syscall.SIGTERM:
		log.Info("Received SIGINT, stopping...")
	}
}
