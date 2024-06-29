// wire.go
//go:build wireinject

package main

import (
	handlers2 "fiap-tech-challenge-producao/internal/adapters/handlers"
	"fiap-tech-challenge-producao/internal/adapters/handlers/http"
	"fiap-tech-challenge-producao/internal/adapters/handlers/pubsub"
	"fiap-tech-challenge-producao/internal/adapters/repository"
	"fiap-tech-challenge-producao/internal/core/usecase"
	db "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/db/mysql"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/messaging"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/middlewares/auth"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/util"

	"github.com/google/wire"
)

func InitializeWebServer() (*handlers2.Server, error) {
	wire.Build(db.NewMySQLConnector,
		util.NewCustomValidator,
		repository.NewFilaRepo,
		pubsub.NewProducaoHandler,
		messaging.NewSqsClient,
		auth.NewJwtToken,
		usecase.NewCadastraFila,
		usecase.NewAtualizaStatusProducaoUC,
		usecase.NewPegaPedidoPorID,
		http.NewHealthCheck,
		http.NewProducao,
		handlers2.NewAPIServer)
	return &handlers2.Server{}, nil
}
