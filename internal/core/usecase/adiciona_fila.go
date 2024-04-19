package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type CadastrarFila interface {
	Cadastra(ctx context.Context, fila *domain.Fila) error
}

type cadastraFila struct {
	filaRepo repository.FilaRepo
}

func (uc *cadastraFila) Cadastra(ctx context.Context, fila *domain.Fila) error {
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
