package usecase

import (
	"context"
	"fiap-tech-challenge-producao/internal/adapters/repository"
	"fmt"
	_error "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/errors"
)

type AtualizaStatusProducao interface {
	Atualiza(ctx context.Context, status, pedidoID string) error
}

type atualizaStatusProducao struct {
	filaRepo repository.FilaRepo
}

func (uc atualizaStatusProducao) Atualiza(ctx context.Context, status, pedidoID string) error {
	ped, err := uc.filaRepo.PegaPedidoPorID(ctx, pedidoID)
	if err != nil {
		return err
	}

	if ped == nil {
		return _error.NotFound.New(fmt.Sprintf("nenhum pedido encontrado para id %s", pedidoID))
	}

	return uc.filaRepo.AtualizaStatus(ctx, status, pedidoID)
}

func NewAtualizaStatusProducaoUC(filaRepo repository.FilaRepo) AtualizaStatusProducao {
	return &atualizaStatusProducao{
		filaRepo: filaRepo,
	}
}
