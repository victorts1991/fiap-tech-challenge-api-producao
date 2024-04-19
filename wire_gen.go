// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fiap-tech-challenge-api/internal/adapters/http"
	"fiap-tech-challenge-api/internal/adapters/http/handlers"
	"fiap-tech-challenge-api/internal/adapters/http/middlewares/auth"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/usecase"
	"fiap-tech-challenge-api/internal/core/usecase/mapper"
	"fiap-tech-challenge-api/internal/util"
)

// Injectors from wire.go:

func InitializeWebServer() (*http.Server, error) {
	healthCheck := handlers.NewHealthCheck()
	dbConnector := repository.NewMySQLConnector()
	validator := util.NewCustomValidator()
	token := auth.NewJwtToken()	
	pedidoRepo := repository.NewPedidoRepo(dbConnector)
	pedido := mapper.NewPedidoMapper()
	listarPedidoPorStatus := usecase.NewListaPedidoPorStatus(pedidoRepo, pedido)
	listarTodosPedidos := usecase.NewListaTodosPedidos(pedidoRepo, pedido)	
	filaRepo := repository.NewFilaRepo(dbConnector)	
	atualizaStatusPedidoUC := usecase.NewAtualizaStatusPedidoUC(pedidoRepo, filaRepo)
	pegarDetalhePedido := usecase.NewPegaDetalhePedido(pedidoRepo, pedido)	
	cadastrarFila := usecase.NewCadastraFila(filaRepo)	
	handlersPedido := handlers.NewPedido(validator, listarPedidoPorStatus, listarTodosPedidos, atualizaStatusPedidoUC, pegarDetalhePedido, cadastrarFila, token)	
	server := http.NewAPIServer(healthCheck, handlersPedido)
	return server, nil
}