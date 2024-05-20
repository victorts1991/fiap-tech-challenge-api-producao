package usecase

import (
	"context"
	"fiap-tech-challenge-producao/internal/adapters/repository"
	"fiap-tech-challenge-producao/internal/core/domain"
)

type CadastrarFila interface {
	Cadastra(ctx context.Context, fila *domain.Producao) error
}

type cadastraFila struct {
	filaRepo repository.FilaRepo
}

func (uc *cadastraFila) Cadastra(ctx context.Context, fila *domain.Producao) error {
	err := uc.filaRepo.Insere(ctx, fila)

	if err != nil {
		return err
	}
	return nil
}

func NewCadastraFila(filaRepo repository.FilaRepo) CadastrarFila {
	return &cadastraFila{
		filaRepo: filaRepo,
	}
}
