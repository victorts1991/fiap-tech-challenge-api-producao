// wire.go
//go:build wireinject

package main

import (
	"fiap-tech-challenge-api/internal/adapters/http"
	"fiap-tech-challenge-api/internal/adapters/http/handlers"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/usecase"
	db "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/db/mysql"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/middlewares/auth"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/util"

	"github.com/google/wire"
)

func InitializeWebServer() (*http.Server, error) {
	wire.Build(db.NewMySQLConnector,
		util.NewCustomValidator,
		repository.NewFilaRepo,
		auth.NewJwtToken,
		usecase.NewCadastraFila,
		usecase.NewAtualizaStatusProducaoUC,
		usecase.NewPegaPedidoPorID,
		handlers.NewHealthCheck,
		handlers.NewProducao,
		http.NewAPIServer)
	return &http.Server{}, nil
}
