// wire.go
//go:build wireinject

package main

import (
	"fiap-tech-challenge-api/internal/adapters/http"
	"fiap-tech-challenge-api/internal/adapters/http/handlers"
	"fiap-tech-challenge-api/internal/adapters/http/middlewares/auth"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/usecase"
	"fiap-tech-challenge-api/internal/core/usecase/mapper"
	"fiap-tech-challenge-api/internal/util"

	"github.com/google/wire"
)

func InitializeWebServer() (*http.Server, error) {
	wire.Build(repository.NewMySQLConnector,
		util.NewCustomValidator,
		repository.NewClienteRepo,
		repository.NewProdutoRepo,
		repository.NewPedidoRepo,
		repository.NewPedidoProdutoRepo,
		repository.NewPagamentoRepo,
		repository.NewFilaRepo,
		auth.NewJwtToken,
		mapper.NewPedidoMapper,
		usecase.NewCadastraCliente,
		usecase.NewPesquisarClientePorCpf,
		usecase.NewCadastraProduto,
		usecase.NewCadastraFila,
		usecase.NewPegaProdutoPorCategoria,
		usecase.NewApagaProduto,
		usecase.NewAtualizaProduto,
		usecase.NewListaPedidoPorStatus,
		usecase.NewListaTodosPedidos,
		usecase.NewCadastraPedido,
		usecase.NewAtualizaStatusPedidoUC,
		usecase.NewPegaDetalhePedido,
		usecase.NewRealizaCheckout,
		usecase.NewPesquisaPagamento,
		handlers.NewCliente,
		handlers.NewProduto,
		handlers.NewHealthCheck,
		handlers.NewPedido,
		handlers.NewPagamento,
		handlers.NewLogin,
		http.NewAPIServer)
	return &http.Server{}, nil
}
