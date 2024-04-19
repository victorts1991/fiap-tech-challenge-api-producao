package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecase/mapper"
	"fmt"
	"sort"
	"strings"
)

var validStatusesSet = map[string]bool{
	domain.StatusEmPreparacao: true,
	domain.StatusFinalizado:   true,
	domain.StatusPronto:       true,
	domain.StatusRecebido:     true,
}

type ListarPedidos interface {
	lista(ctx context.Context) ([]*domain.Pedido, error)
}

type ListarPedidoPorStatus interface {
	ListaPorStatus(ctx context.Context, statuses []string) ([]*domain.Pedido, error)
}

type listaPedido struct {
	repo         repository.PedidoRepo
	mapperPedido mapper.Pedido
}

type ListarTodosPedidos interface {
	ListaTodos(ctx context.Context) ([]*domain.Pedido, error)
}

func (uc *listaPedido) ListaPorStatus(ctx context.Context, statuses []string) ([]*domain.Pedido, error) {
	if err := ValidaStatuses(statuses); err != nil {
		return nil, err
	}

	pedidos, err := uc.repo.PesquisaPorStatus(ctx, statuses)
	if err != nil {
		return nil, err
	}
	return uc.mapperPedido.MapDTOToModels(pedidos), nil
}

func NewListaPedidoPorStatus(repo repository.PedidoRepo, mapperPedido mapper.Pedido) ListarPedidoPorStatus {
	return &listaPedido{
		repo:         repo,
		mapperPedido: mapperPedido,
	}
}

func ValidaStatuses(statuses []string) error {
	for _, s := range statuses {
		if !validStatusesSet[strings.ToLower(s)] {
			return commons.BadRequest.New(fmt.Sprintf("%s não é um status valido", s))
		}
	}

	return nil
}

func (uc *listaPedido) ListaTodos(ctx context.Context) ([]*domain.Pedido, error) {
	pedidos, err := uc.repo.PesquisaTodos(ctx)

	if err != nil {
		return nil, err
	}

	sort.Sort(domain.PedidosDTO(pedidos))

	return uc.mapperPedido.MapDTOToModels(pedidos), nil
}

func NewListaTodosPedidos(repo repository.PedidoRepo, mapperPedido mapper.Pedido) ListarTodosPedidos {
	return &listaPedido{
		repo:         repo,
		mapperPedido: mapperPedido,
	}
}
