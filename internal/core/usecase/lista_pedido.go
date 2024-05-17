package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type PegaPedidoPorID interface {
	PesquisaPorID(ctx context.Context, pedidoID string) (*domain.Producao, error)
}

type listaPedidoPorID struct {
	repo repository.FilaRepo
}

func (uc *listaPedidoPorID) PesquisaPorID(ctx context.Context, pedidoID string) (*domain.Producao, error) {
	return uc.repo.PegaPedidoPorID(ctx, pedidoID)
}

func NewPegaPedidoPorID(repo repository.FilaRepo) PegaPedidoPorID {
	return &listaPedidoPorID{
		repo: repo,
	}
}
