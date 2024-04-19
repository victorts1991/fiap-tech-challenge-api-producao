package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fmt"
)

type AtualizaStatusPedidoUC interface {
	Atualiza(ctx context.Context, status string, id int64) error
}

type atualizaStatusPedido struct {
	repo     repository.PedidoRepo
	filaRepo repository.FilaRepo
}

func (uc atualizaStatusPedido) Atualiza(ctx context.Context, status string, id int64) error {
	if err := ValidaStatuses([]string{status}); err != nil {
		return err
	}

	ped, err := uc.repo.PesquisaPorID(ctx, id)
	if err != nil {
		return err
	}

	if couldNotUpdateStatus(ped.Status) {
		return commons.BadRequest.New(fmt.Sprintf("pedido %d não pode atualizar para %s, status atual: %s", ped.Id, status, ped.Status))
	}

	if err = uc.filaRepo.AtualizaStatus(ctx, status, id); err != nil {
		return err
	}

	return uc.repo.AtualizaStatus(ctx, status, id)
}

func couldNotUpdateStatus(status string) bool {
	return status == domain.StatusAguardandoPagamento ||
		status == domain.StatusPagamentoRecusado
}

func NewAtualizaStatusPedidoUC(repo repository.PedidoRepo, filaRepo repository.FilaRepo) AtualizaStatusPedidoUC {
	return &atualizaStatusPedido{
		repo:     repo,
		filaRepo: filaRepo,
	}
}
